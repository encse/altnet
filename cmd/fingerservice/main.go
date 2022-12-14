package main

import (
	"io/ioutil"
	"net"
	"strings"

	"github.com/encse/altnet/lib/csokavar"
	"github.com/encse/altnet/lib/io"
	log "github.com/sirupsen/logrus"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func main() {
	l, err := net.Listen("tcp", ":8000")
	io.FatalIfError(err)
	defer func() {
		log.Info("fingerservice exit")
		l.Close()
	}()
	host, port, err := net.SplitHostPort(l.Addr().String())
	io.FatalIfError(err)

	log.Info("Finger service listening on host: %s, port: %s\n", host, port)

	for {
		conn, err := l.Accept()
		io.FatalIfError(err)

		go func(conn net.Conn) {
			buf := make([]byte, 1024)
			len, err := conn.Read(buf)
			if err != nil {
				log.Errorf("Error reading: %#v", err)
				return
			}
			user := strings.TrimSpace(string(buf[:len]))
			message := asciiFold(csokavar.Finger(user, 80))
			conn.Write([]byte(message))
			conn.Close()
		}(conn)
	}
}

func asciiFold(st string) string {
	st = strings.ReplaceAll(st, "█", "#")
	st = strings.ReplaceAll(st, "▀", "\"")
	st = strings.ReplaceAll(st, "▌", ";")
	st = strings.ReplaceAll(st, "▐", ":")
	st = strings.ReplaceAll(st, "▄", ".")

	// remove accents such as á -> a, é -> e, because raw TCP doesn't like it...
	isNotAscii := func(r rune) bool {
		return r >= 127
	}
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isNotAscii), norm.NFC)
	data, _ := ioutil.ReadAll(transform.NewReader(strings.NewReader(st), t))
	return string(data)
}
