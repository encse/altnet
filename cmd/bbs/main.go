package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/encse/altnet/lib/config"
	"github.com/encse/altnet/lib/csokavar"
	"github.com/encse/altnet/lib/io"
	log "github.com/sirupsen/logrus"
	"golang.org/x/term"
)

func main() {

	conf := config.Get()

	screenWidth, _, err := term.GetSize(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(csokavar.Banner(screenWidth))

	fmt.Println("Enter your username or GUEST")

	username := ""
	for username == "" {
		username, err = io.Readline("Username: ")
		if err != nil {
			log.Fatal(err)
		}
		username = strings.ToLower(username)
	}

	if username != "guest" {
		for i := 0; i < 3; i++ {
			fmt.Print("Password: ")
			_, err = io.ReadPassword()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("")
		}
		return
	}

	logo, err := csokavar.Logo(screenWidth)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(logo)
	fmt.Println("Welcome", username)

loop:
	for {
		fmt.Println("BBS Menu")
		fmt.Println("------------")
		options := ""
		fmt.Println(": Latest [T]weets")
		options += "t"
		fmt.Println(": [G]itHub skyline")
		options += "g"
		fmt.Println(": [C]ontact sysop")
		options += "c"
		if conf.Dfrotz.Location != "" {
			fmt.Println(": play [I]dőrégész")
			options += "i"
		}
		fmt.Println(": e[X]it")
		options += "x"

		option, err := io.ReadOption("Select an item", options)
		if err != nil {
			log.Fatal(err)
		}
		switch strings.ToLower(option) {
		case "t":
			tweets, err := csokavar.GetTweets("encse", screenWidth)
			if err != nil {
				log.Error(err)
				tweets = "Could not get tweets now."
			}
			fmt.Println(tweets)
		case "g":
			skyline, err := csokavar.GetSkyline("encse", screenWidth)
			if err != nil {
				log.Error(err)
				skyline = "Could not get skyline now."
			}
			fmt.Println(skyline)
		case "c":
			gpgKey, err := csokavar.GpgKey(screenWidth)
			if err != nil {
				log.Error(err)
				gpgKey = "Could not get contact info now."
			}
			fmt.Println(gpgKey)
		case "i":
			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt)
			go func() {
				for range c {
					// pass
				}
			}()
			cmd := exec.Command(conf.Dfrotz.Location, "-r", "lt", "-R", "/tmp", "data/doors/idoregesz.z5")
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
			signal.Stop(c)
			if err != nil {
				log.Error(err)
				fmt.Println("An error occured.")
			}
		case "x":
			break loop
		}
	}

	fmt.Println("Have a nice day!")

	footer, err := csokavar.Footer(screenWidth)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(footer)
}
