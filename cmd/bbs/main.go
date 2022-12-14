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
	io.FatalIfError(err)

	fmt.Println(csokavar.Banner(screenWidth))
	fmt.Println("Enter your username or GUEST")

	username, err := io.ReadNotEmpty("Username: ")
	io.FatalIfError(err)

	username = strings.ToLower(username)

	if username != "guest" {
		for i := 0; i < 3; i++ {
			_, err = io.ReadPassword("Password: ")
			io.FatalIfError(err)
		}
		return
	}

	logo, err := csokavar.Logo(screenWidth)
	io.FatalIfError(err)

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
		fmt.Println(": [s]hell")
		options += "s"
		fmt.Println(": e[X]it")
		options += "x"

		option, err := io.ReadOption("Select an item", options)
		io.FatalIfError(err)

		switch strings.ToLower(option) {
		case "t":
			runCommand("./twitter", "encse")
		case "g":
			runCommand("./skyline", "encse")
		case "c":
			gpgKey, err := csokavar.GpgKey(screenWidth)
			if err != nil {
				log.Error(err)
				gpgKey = "Could not get contact info now."
			}
			fmt.Println(gpgKey)
		case "i":
			runCommand(conf.Dfrotz.Location, "-r", "lt", "-R", "/tmp", "data/doors/idoregesz.z5")
		case "s":
			runCommand("./shell")
		case "x":
			break loop
		}
	}

	fmt.Println("Have a nice day!")

	footer, err := csokavar.Footer(screenWidth)
	if err != nil {
		log.Error(err)
		return
	}
	fmt.Println(footer)
}

func runCommand(name string, arg ...string) {
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
}
