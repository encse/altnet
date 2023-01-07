package main

import (
	"fmt"
	"strings"

	"github.com/encse/altnet/lib/csokavar"
	"github.com/encse/altnet/lib/io"
)

func main() {
	for {
		cmd, err := io.ReadNotEmpty("> ")
		io.FatalIfError(err)

		cmd = strings.TrimSpace(cmd)
		parts := strings.Split(cmd, " ")
		if len(parts) > 0 {
			switch parts[0] {
			case "?", "help":
				fmt.Println("uumap <host>\tshow uumap entry for a host")
				fmt.Println("uuplot <host>\tplot uupath to a host")
				fmt.Println("twitter <user>\tshow the latest tweets of a Twitter user")
				fmt.Println("skyline <user> \tshow the GitHub contributions chart for a GitHub user")
				fmt.Println("zrun <game>\tplay Z-machine games")
				break
			case "zrun":
				csokavar.RunCommand("./zrun", parts[1:]...)
			case "uumap":
				csokavar.RunCommand("./uumap", parts[1:]...)
			case "uuplot":
				csokavar.RunCommand("./uuplot", parts[1:]...)
			case "twitter":
				csokavar.RunCommand("./twitter", parts[1:]...)
			case "skyline":
				csokavar.RunCommand("./skyline", parts[1:]...)
			case "exit", "quit":
				return
			default:
				fmt.Println("Unknown command. For the list of commands enter ?; or help. ")
			}
		}
	}
}
