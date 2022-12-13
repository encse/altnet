package main

import (
	"fmt"

	"github.com/encse/altnet/lib/csokavar"
	log "github.com/sirupsen/logrus"
)

func main() {

	st, err := csokavar.GetSkyline("encse", 120)
	if err != nil {
		log.Error(err)
		return
	}
	fmt.Println(st)
}
