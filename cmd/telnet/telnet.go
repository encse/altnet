package main

import (
	"context"
	"fmt"
	"os"

	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/slices"
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

	host, err := network.Lookup(ctx, hostName)
	io.FatalIfError(err)
	if host == nil {
		fmt.Println("host not found")
		return
	}

	targetHostName := schema.HostName("")
	if len(os.Args) > 1 {
		targetHostName = schema.HostName(os.Args[1])
	} else {
		targetHostName, err = io.ReadArgFromList("host", os.Args, 1, host.Neighbours)
		io.FatalIfError(err)
	}

	if !slices.Contains(host.Neighbours, targetHostName) {
		fmt.Println("host not in NETSTAT")
		return
	}

	targetHost, err := network.Lookup(ctx, targetHostName)
	io.FatalIfError(err)

	altnet.Login(ctx, targetHost)

	fmt.Println("Connection closed")
}
