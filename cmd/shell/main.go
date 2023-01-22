package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/csokavar"
	"github.com/encse/altnet/lib/io"
)

func main() {
	ctx := altnet.ContextFromEnv(context.Background())
	host, err := altnet.GetHost(ctx)
	io.FatalIfError(err)
	for {
		cmd, err := io.ReadNotEmpty(fmt.Sprintf("%v$ ", host))
		io.FatalIfError(err)

		cmd = strings.TrimSpace(cmd)
		parts := strings.Split(cmd, " ")
		if len(parts) > 0 {
			switch parts[0] {
			case "?", "help":
				fmt.Println(io.Table(
					[]string{"cat <file>", "print file contents"},
					[]string{"hosts", "show a list of hosts on the network"},
					[]string{"ls <glob>", "list files"},
					[]string{"netstat", "show connected hosts"},
					[]string{"ps ", "show running processes on this host"},
					[]string{"skyline <user>", "show the GitHub contributions chart for a GitHub user"},
					[]string{"telnet <host>", "connect to a host in netstat"},
					[]string{"twitter <user>", "show the latest tweets of a Twitter user"},
					[]string{"uumap <host>", "show uumap entry for a host"},
					[]string{"uuplot <host>", "plot uupath to a host"},
					[]string{"zrun <game>", "play Z-machine games"},
				))
				break
			case "hosts":
				csokavar.RunCommand(ctx, "./hosts", parts[1:]...)
			case "telnet":
				csokavar.RunCommand(ctx, "./telnet", parts[1:]...)
			case "netstat":
				csokavar.RunCommand(ctx, "./netstat", parts[1:]...)
			case "ps":
				csokavar.RunCommand(ctx, "./ps", parts[1:]...)
			case "zrun":
				csokavar.RunCommand(ctx, "./zrun", parts[1:]...)
			case "uumap":
				csokavar.RunCommand(ctx, "./uumap", parts[1:]...)
			case "uuplot":
				csokavar.RunCommand(ctx, "./uuplot", parts[1:]...)
			case "twitter":
				csokavar.RunCommand(ctx, "./twitter", parts[1:]...)
			case "skyline":
				csokavar.RunCommand(ctx, "./skyline", parts[1:]...)
			case "ls":
				csokavar.RunCommand(ctx, "./ls", parts[1:]...)
			case "cat":
				csokavar.RunCommand(ctx, "./cat", parts[1:]...)
			case "exit", "quit":
				return
			default:
				fmt.Println("Unknown command. For the list of commands enter ?; or help. ")
			}
		}
	}
}
