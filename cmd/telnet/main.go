package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/csokavar"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/log"
	"github.com/encse/altnet/lib/uumap"
)

func main() {
	ctx := altnet.ContextFromEnv(context.Background())
	host, err := altnet.GetHost(ctx)
	io.FatalIfError(err)

	uumap, err := uumap.GetUumap()
	io.FatalIfError(err)

	entry, ok := uumap[string(host)]
	if !ok {
		fmt.Println("host not found")
		return
	}

	targetHost, err := io.ReadArgFromList("host", os.Args, 1, entry.Hosts)
	io.FatalIfError(err)

	fmt.Println(fmt.Sprintf("Connected to %s", strings.ToUpper(string(targetHost))))
	fmt.Println()
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
	ctx = altnet.SetHost(ctx, altnet.Host(targetHost))
	ctx = altnet.SetUser(ctx, altnet.User(username))

	log.Infof("Connected as %s", username)
	fmt.Println("Welcome", username)
	csokavar.RunHiddenCommand(ctx, "./shell")
	fmt.Println("Connection closed")
}
