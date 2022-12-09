package main

import (
	"fmt"

	"github.com/encse/altnet/lib/config"
	"github.com/encse/altnet/lib/twitter"
)

func main() {
	config, err := config.Read("../../config.yml")
	if err != nil {
		fmt.Println(err)
		return
	}

	st, err := twitter.GetTweets(config.Twitter.AccessToken, "encse", 80)
	if err == nil {
		fmt.Println(st)
	} else {
		fmt.Println(err)
	}
}
