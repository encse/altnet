package main

import (
	"context"
	"fmt"
	"strings"
	"syscall"
	"time"

	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/config"
	"github.com/encse/altnet/lib/csokavar"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/log"
	"github.com/encse/altnet/lib/uman"
	"github.com/encse/altnet/lib/uumap"
	"github.com/encse/altnet/schema"
	"golang.org/x/term"
)

func main() {
	ctx := altnet.ContextFromEnv(context.Background())
	conf := config.Get()

	network, err := uumap.NetworkConn()
	io.FatalIfError(err)
	defer network.Close()

	screenWidth, _, err := term.GetSize(int(syscall.Stdin))
	io.FatalIfError(err)

	fmt.Println(csokavar.Banner(screenWidth))

	loginres, err := login(ctx, network)
	io.FatalIfError(err)
	if loginres == nil {
		return
	}

	ctx = altnet.SetRealUser(ctx, loginres.User)
	log.Infof("Connected as %s", loginres.User)

	ctx = altnet.SetUser(ctx, loginres.User)
	fmt.Println()
	fmt.Println("Welcome", loginres.User)
	if loginres.LastLoginFailure != nil {
		fmt.Println("Last login failure", loginres.LastLoginFailure.Format(time.RFC1123))
	}

	if loginres.LastLogin != nil {
		fmt.Println("Last login", loginres.LastLogin.Format(time.RFC1123))
	}
	fmt.Println()

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
			altnet.RunCommand(ctx, "./twitter", "encse")
		case "g":
			altnet.RunCommand(ctx, "./skyline", "encse")
		case "c":
			gpgKey, err := csokavar.GpgKey(screenWidth)
			if err != nil {
				log.Error(err)
				gpgKey = "Could not get contact info now."
			}
			fmt.Println(gpgKey)
		case "i":
			altnet.RunCommand(ctx, "./zrun", "idoregesz")
		case "s":
			altnet.RunHiddenCommand(ctx, "./shell")
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

func login(ctx context.Context, network uumap.Network) (*uman.LoginRes, error) {
	fmt.Println("Enter your username or GUEST. If you don't have a user yet, enter NEWUSER.")

	username, err := io.ReadNotEmpty[schema.Uname]("Username: ")
	io.FatalIfError(err)

	username = schema.Uname(strings.ToLower(string(username)))
	if username == "guest" {
		return &uman.LoginRes{User: username}, nil
	} else if username == "newuser" {
		return uman.RegisterUser(ctx, network)
	} else {
		for i := 0; i < 3; i++ {
			password, err := io.ReadPassword("Password: ")
			io.FatalIfError(err)

			loginres, err := uman.LoginAttempt(ctx, network, username, password)
			if loginres != nil || err != nil {
				return loginres, err
			}
		}
	}
	return nil, nil
}
