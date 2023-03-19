package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/log"
	"github.com/encse/altnet/lib/uumap"
	"github.com/encse/altnet/schema"
)

func main() {
	ctx := altnet.ContextFromEnv(context.Background())
	hostName, err := altnet.GetHost(ctx)
	io.FatalIfError(err)

	network, err := uumap.NetworkConn()
	io.FatalIfError(err)
	defer network.Close()

	fmt.Printf("Connected to %s\n", strings.ToUpper(string(hostName)))
	fmt.Println()
	fmt.Println("Enter your username or GUEST")

	host, err := network.Lookup(ctx, hostName)
	io.FatalIfError(err)

	username, err := io.ReadNotEmpty[schema.Uname]("Login: ")
	io.FatalIfError(err)

	username = username.ToLower()

	valid := username == "guest"
	for i := 0; i < 3 && !valid; i++ {
		password, err := io.ReadPassword("Password: ")
		valid, err = altnet.ValidatePassword(
			ctx, network, host, username, password,
		)
		io.FatalIfError(err)
	}

	if !valid {
		return
	}

	ctx = altnet.SetUser(ctx, username)
	log.Infof("Connected as %s", username)
	fmt.Println("Welcome", username)
	altnet.RunHiddenCommand(ctx, "./shell")
}
