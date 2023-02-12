package main

import (
	"context"
	"net/http"
	"sync/atomic"

	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/csokavar"
	ioutil "github.com/encse/altnet/lib/io"
	log "github.com/encse/altnet/lib/log"
)

func main() {
	var sessionCounter uint32

	http.Handle("/", http.FileServer(http.Dir("data/www")))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		sessionId := altnet.SessionId(atomic.AddUint32(&sessionCounter, 1))

		session, err := NewSession(sessionId, w, r)
		if err != nil {
			log.Error(err)
			return
		}

		err = session.Run(context.Background())
		if err != nil {
			log.Error(err)
		}

		err = session.Close()
		if err != nil {
			log.Error(err)
		}
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
