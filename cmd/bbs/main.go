package main

import (
	"context"
	"fmt"
	"strings"
	"syscall"

	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/config"
	"github.com/encse/altnet/lib/csokavar"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/log"
	"golang.org/x/term"
)

func main() {
	ctx := altnet.ContextFromEnv(context.Background())
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
	ctx = altnet.SetUser(ctx, altnet.User(username))

	log.Infof("Connected as %s", username)
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
			csokavar.RunCommand(ctx, "./twitter", "encse")
		case "g":
			csokavar.RunCommand(ctx, "./skyline", "encse")
		case "c":
			gpgKey, err := csokavar.GpgKey(screenWidth)
			if err != nil {
				log.Error(err)
				gpgKey = "Could not get contact info now."
			}
			fmt.Println(gpgKey)
		case "i":
			csokavar.RunCommand(ctx, "./zrun", "idoregesz")
		case "s":
			csokavar.RunHiddenCommand(ctx, "./shell")
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
