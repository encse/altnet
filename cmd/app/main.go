package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"time"

	"github.com/creack/pty"
	"github.com/encse/altnet/lib/csokavar"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	Subprotocols:    []string{"9600", "19200", "56000"},
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
	send    chan rune
	receive chan byte
	bps     int
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

	for _, b := range []rune(string(p)) {
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
	conn := Conn{
		send:    make(chan rune, 2048),
		receive: make(chan byte, 2048),
		bps:     speed / 8,
	}

	c.SetReadLimit(maxMessageSize)
	c.SetReadDeadline(time.Now().Add(pongWait))
	c.SetPongHandler(func(string) error { c.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	go func() {
		defer func() {
			fmt.Println("exit reader")
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
			fmt.Println("exit writer")
			ticker.Stop()
			c.Close()
		}()

		for {
			select {
			case message, ok := <-conn.send:
				c.SetWriteDeadline(time.Now().Add(writeWait))
				if !ok {
					c.WriteMessage(websocket.CloseMessage, []byte{})
					return
				}

				err := c.WriteMessage(websocket.TextMessage, []byte(string(message)))
				if err != nil {
					fmt.Println(err)
					return
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
		shell(conn, conn)
		conn.Close()
	}()
}

func shell(writer io.Writer, reader io.Reader) error {
	// Create arbitrary command.
	c := exec.Command("./bbs")

	// Start the command with a pty.
	ptmx, err := pty.StartWithSize(c, &pty.Winsize{Cols: 80, Rows: 25})
	if err != nil {
		return err
	}
	// Make sure to close the pty at the end.
	defer func() { _ = ptmx.Close() }() // Best effort.

	go func() { _, _ = io.Copy(ptmx, reader) }()
	_, _ = io.Copy(writer, ptmx)

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
		fmt.Println(st)
		w.Write([]byte(st))
	})

	log.Print("listening on port 7979")
	if err := http.ListenAndServe(":7979", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
