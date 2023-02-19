package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/slices"
	"github.com/encse/altnet/lib/uumap"
)

func main() {
	ctx := altnet.ContextFromEnv(context.Background())
	host, err := altnet.GetHost(ctx)
	io.FatalIfError(err)

	entries, err := uumap.GetUumap()
	io.FatalIfError(err)

	entry, ok := entries[string(host)]
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

	altnet.Login(ctx, uumap.Host(targetHost))

	fmt.Println("Connection closed")
}
