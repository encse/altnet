package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"syscall"
	"time"

	"github.com/creack/pty"
	"github.com/encse/altnet/lib/altnet"
	log "github.com/encse/altnet/lib/log"
	"github.com/gorilla/websocket"
	errgroup "golang.org/x/sync/errgroup"
)

const (
	writeWait      = 10 * time.Second    // Time allowed to write a message to the peer.
	pongWait       = 60 * time.Second    // Time allowed to read the next pong message from the peer.
	pingPeriod     = (pongWait * 9) / 10 // Send pings to peer with this period. Must be less than pongWait.
	maxMessageSize = 512                 // Maximum message size allowed from peer.
)

type Session struct {
	sessionId     altnet.SessionId
	outputBuffer  chan byte
	inputBuffer   chan byte
	bps           int
	connectedFrom string
	ws            *websocket.Conn
}

func NewSession(sessionId altnet.SessionId, w http.ResponseWriter, r *http.Request) (conn Session, err error) {

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		Subprotocols:    []string{"2400", "4800", "9600", "14400", "19200", "28800", "33600", "56000"},
	}

	wsconn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return Session{}, err
	}

	wsconn.SetReadLimit(maxMessageSize)

	err = wsconn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		err = errors.Join(err, wsconn.Close())
		return Session{}, err
	}
	wsconn.SetPongHandler(func(string) error {
		return wsconn.SetReadDeadline(time.Now().Add(pongWait))
	})

	s := Session{
		sessionId:     sessionId,
		outputBuffer:  make(chan byte, 2048),
		inputBuffer:   make(chan byte, 2048),
		bps:           getSpeed(wsconn.Subprotocol()) / 8,
		connectedFrom: getConnectedFrom(r),
		ws:            wsconn,
	}
	return s, nil
}

func (s Session) Close() error {
	return s.ws.Close()
}

func (session Session) Run(ctx context.Context) error {
	log.Infof("enter session")
	defer log.Info("exit session")

	cmd := exec.Command("./bbs")
	cmd.Env = os.Environ()
	cmd.Env = append(os.Environ(),
		fmt.Sprintf("CONNECTED_FROM=%v", session.connectedFrom),
		fmt.Sprintf("ALTNET_HOST=%v", "csokavar"),
		fmt.Sprintf("ALTNET_SESSION=%v", session.sessionId),
		fmt.Sprintf("TERM=%v", "xterm-color"),
	)

	cmd.Stderr = os.Stderr
	ptmx, err := pty.StartWithSize(cmd, &pty.Winsize{Cols: 80, Rows: 25})
	if err != nil {
		return err
	}

	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error { return pingLoop(ctx, session) })
	group.Go(func() error { return inputLoop(ctx, ptmx, session) })
	group.Go(func() error { return outputLoop(ctx, ptmx, session) })
	group.Go(func() error { return waitLoop(ctx, cmd, ptmx, session) })

	err = group.Wait()
	if err == io.EOF {
		err = nil
	}
	return err
}

func getSpeed(speedString string) int {
	speed, err := strconv.Atoi(speedString)
	if err != nil {
		log.Error(err)
		return 560000
	}
	return speed
}

func getConnectedFrom(r *http.Request) string {
	connectedFrom := ""
	if items := r.Header["X-Real-Ip"]; len(items) > 0 {
		connectedFrom = items[0]
	}
	log.Info("Connected from  ", connectedFrom)
	return connectedFrom
}

func waitLoop(ctx context.Context, cmd *exec.Cmd, ptmx *os.File, s Session) error {
	log.Info("enter wait")
	defer log.Info("exit wait")
	<-ctx.Done()

	err := ptmx.Close()
	if err != nil {
		log.Error(err)
	}

	err = altnet.KillSession(s.sessionId, syscall.SIGHUP)
	if err != nil {
		log.Error(err)
	}

	err = cmd.Wait()
	if err != nil {
		log.Error(err)
	}
	return err
}

func pingLoop(ctx context.Context, conn Session) error {
	log.Info("enter ping loop")
	defer log.Info("exit ping loop")

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(pingPeriod):
			conn.ws.SetWriteDeadline(time.Now().Add(writeWait))
			err := conn.ws.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				return err
			}
		}
	}
}

func inputLoop(ctx context.Context, ptmx io.Writer, session Session) error {
	log.Info("enter input loop")
	defer log.Info("exit input loop")

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			_, message, err := session.ws.ReadMessage()
			if err != nil {
				return err
			}

			_, err = ptmx.Write(message)
			if err != nil {
				return err
			}
		}
	}
}

func outputLoop(ctx context.Context, ptmx io.Reader, session Session) error {
	log.Info("enter output loop")
	defer log.Info("exit output loop")

	sent := 0
	started := time.Now()
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			buffer := make([]byte, 1)
			_, err := ptmx.Read(buffer)

			if err != nil {
				_ = session.ws.SetWriteDeadline(time.Now().Add(writeWait))
				return session.ws.WriteMessage(websocket.CloseMessage, []byte{})
			}

			if sent == session.bps/10 {
				timeSpent := time.Since(started)
				if timeSpent < 100*time.Millisecond {
					time.Sleep(100*time.Millisecond - timeSpent)
				}
				sent = 0
			}

			if sent == 0 {
				started = time.Now()
			}

			_ = session.ws.SetWriteDeadline(time.Now().Add(writeWait))
			err = session.ws.WriteMessage(websocket.BinaryMessage, buffer)
			if err != nil {
				return err
			}
			sent++
		}
	}
}
