package main

import (
	"fmt"

	"github.com/encse/altnet/lib/csokavar"
)

func main() {
	st, err := csokavar.GetTweets("encse", 80)
	if err == nil {
		fmt.Println(st)
	} else {
		fmt.Println(err)
	}
}
