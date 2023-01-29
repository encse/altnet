package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/csokavar"
	"github.com/encse/altnet/lib/io"
)

var commands = [][]string{
	{"ansi <file>", "ansi file (*.ans) viewer"},
	{"cat <file>", "print file contents"},
	{"hosts", "show a list of hosts on the network"},
	{"ls <glob>", "list files"},
	{"netstat", "show connected hosts"},
	{"ps ", "show running processes on this host"},
	{"skyline <user>", "show the GitHub contributions chart for a GitHub user"},
	{"telnet <host>", "connect to a host in netstat"},
	{"twitter <user>", "show the latest tweets of a Twitter user"},
	{"uumap <host>", "show uumap entry for a host"},
	{"uuplot <host>", "plot uupath to a host"},
	{"zrun <game>", "play Z-machine games"},
}

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
			exe := getExe(parts[0], commands)
			if exe != "" {
				csokavar.RunCommand(ctx, exe, parts[1:]...)
			} else if parts[0] == "?" {
				fmt.Println(io.Table(commands...))
			} else if parts[0] == "help" {
				fmt.Println(io.Table(commands...))
			} else if parts[0] == "exit" {
				return
			} else if parts[0] == "quit" {
				return
			} else {
				fmt.Println("Unknown command. For the list of commands enter ?; or help. ")
			}
		}
	}
}

func getExe(cmd string, commands [][]string) string {
	for _, item := range commands {
		cmdName := strings.Split(item[0], " ")[0]
		if cmd == cmdName {
			return "./" + cmdName
		}
	}
	return ""
}
