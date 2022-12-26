package main

import (
	"errors"
	"fmt"
	"io"
	"unicode/utf8"

	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/creack/pty"
	"github.com/encse/altnet/lib/csokavar"
	ioutil "github.com/encse/altnet/lib/io"
	log "github.com/encse/altnet/lib/log"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	Subprotocols:    []string{"2400", "4800", "9600", "14400", "19200", "28800", "33600", "56000"},
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type Conn struct {
	send          chan byte
	receive       chan byte
	bps           int
	connectedFrom string
}

func (conn Conn) Close() {
	close(conn.send)
}

func (conn Conn) Read(p []byte) (n int, err error) {
	b, ok := <-conn.receive
	if !ok {
		return 0, errors.New("cannot read")
	}
	p[0] = b
	return 1, nil
}

func (conn Conn) Write(p []byte) (n int, err error) {
	sent := 0
	started := time.Now()
	for _, b := range p {
		if sent == conn.bps/10 {
			timeSpent := time.Now().Sub(started)
			if timeSpent < 100*time.Millisecond {
				time.Sleep(100*time.Millisecond - timeSpent)
			}
			sent = 0
		}
		if sent == 0 {
			started = time.Now()
		}
		conn.send <- b
		sent++

	}
	return len(p), nil
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}
	speed, err := strconv.Atoi(c.Subprotocol())

	if err != nil {
		log.Println(err)
		return
	}

	connectedFrom := ""
	if items := r.Header["X-Real-Ip"]; len(items) > 0 {
		connectedFrom = items[0]
	}
	log.Info("Connected from  ", connectedFrom)

	conn := Conn{
		send:          make(chan byte, 2048),
		receive:       make(chan byte, 2048),
		bps:           speed / 8,
		connectedFrom: connectedFrom,
	}

	c.SetReadLimit(maxMessageSize)
	c.SetReadDeadline(time.Now().Add(pongWait))
	c.SetPongHandler(func(string) error { c.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	go func() {
		defer func() {
			log.Info("exit reader")
			c.Close()
		}()
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Printf("error: %v", err)
				break
			}
			for _, b := range message {
				conn.receive <- b
			}
		}
	}()

	go func() {
		ticker := time.NewTicker(pingPeriod)
		defer func() {
			log.Info("exit writer")
			ticker.Stop()
			c.Close()
		}()

		chunk := make([]byte, 0, 4)

		for {
			select {
			case b, ok := <-conn.send:
				chunk = append(chunk, b)
				if utf8.FullRune(chunk) {
					c.SetWriteDeadline(time.Now().Add(writeWait))
					if !ok {
						c.WriteMessage(websocket.CloseMessage, []byte{})
						return
					}
					err := c.WriteMessage(websocket.TextMessage, chunk)
					if err != nil {
						log.Error(err)
						return
					}
					chunk = chunk[:0]
				}

			case <-ticker.C:
				c.SetWriteDeadline(time.Now().Add(writeWait))
				if err := c.WriteMessage(websocket.PingMessage, nil); err != nil {
					return
				}
			}
		}
	}()

	go func() {
		shell(conn)
		conn.Close()
	}()
}

func shell(conn Conn) error {
	// Create arbitrary command.
	c := exec.Command("./bbs")
	c.Env = append(c.Env, fmt.Sprintf("AN_FROM=%s", conn.connectedFrom))
	c.Stderr = os.Stderr
	// Start the command with a pty.
	ptmx, err := pty.StartWithSize(c, &pty.Winsize{Cols: 80, Rows: 25})
	if err != nil {
		return err
	}
	// Make sure to close the pty at the end.
	defer func() { _ = ptmx.Close() }() // Best effort.

	go func() { _, _ = io.Copy(ptmx, conn) }()
	_, _ = io.Copy(conn, ptmx)

	return nil
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("data/www")))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(w, r)
	})

	http.HandleFunc("/~encse", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		st := csokavar.Finger("encse", 120)
		w.Write([]byte(st))
	})

	log.Print("listening on port 7979")
	if err := http.ListenAndServe(":7979", nil); err != nil {
		ioutil.FatalIfError(err)
	}
}
