package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"

	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/log"
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
				break
			case "uumap":
				runCommand("./uumap", parts[1:]...)
			case "uuplot":
				runCommand("./uuplot", parts[1:]...)
			case "twitter":
				runCommand("./twitter", parts[1:]...)
			case "skyline":
				runCommand("./skyline", parts[1:]...)
			case "exit", "quite":
				return
			default:
				fmt.Println("Unknown command. For the list of commands enter ?; or help. ")
			}
		}
	}
}

func runCommand(name string, arg ...string) {
	log.Info("run", name, arg)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer signal.Stop(c)
	go func() {
		for range c {
			// pass
		}
	}()
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Error(err)
	}
	fmt.Println()
}
