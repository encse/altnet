package main

import (
	"context"
	"fmt"
	"os"

	"github.com/encse/altnet/ent/schema"
	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/slices"
	"github.com/encse/altnet/lib/uumap"
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

	targetHost := schema.HostName("")
	if len(os.Args) > 1 {
		targetHost = schema.HostName(os.Args[1])
	} else {
		targetHost, err = io.ReadArgFromList("host", os.Args, 1, host.Neighbours)
		io.FatalIfError(err)
	}

	if !slices.Contains(host.Neighbours, targetHost) {
		fmt.Println("host not in NETSTAT")
		return
	}

	altnet.Login(ctx, host)

	fmt.Println("Connection closed")
}
