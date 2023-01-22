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
	"github.com/encse/altnet/lib/slices"
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

	targetHost := ""
	if len(os.Args) > 1 {
		targetHost = os.Args[1]
	} else {
		targetHost, err = io.ReadArgFromList("host", os.Args, 1, entry.Hosts)
		io.FatalIfError(err)
	}

	targetHost = strings.ToLower(targetHost)

	if !slices.Contains(entry.Hosts, targetHost) {
		fmt.Println("host not in NETSTAT")
		return
	}

	fmt.Println(fmt.Sprintf("Connected to %s", strings.ToUpper(targetHost)))
	fmt.Println()
	fmt.Println("Enter your username or GUEST")

	username, err := io.ReadNotEmpty("Login: ")
	io.FatalIfError(err)

	username = strings.ToLower(username)
	if username != "guest" {
		for i := 0; i < 3; i++ {
			_, err = io.ReadPassword("Password: ")
			io.FatalIfError(err)
		}
	} else {
		ctx = altnet.SetHost(ctx, altnet.Host(targetHost))
		ctx = altnet.SetUser(ctx, altnet.User(username))

		log.Infof("Connected as %s", username)
		fmt.Println("Welcome", username)
		csokavar.RunHiddenCommand(ctx, "./shell")
	}
	fmt.Println("Connection closed")
}
