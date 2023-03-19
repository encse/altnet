package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/uumap"
	"github.com/encse/altnet/schema"
)

func main() {
	ctx := context.Background()

	hostname, err := io.ReadArg("host", os.Args, 1)
	io.FatalIfError(err)

	host := schema.HostName(strings.ToLower(hostname))

	network, err := uumap.NetworkConn()
	io.FatalIfError(err)
	defer network.Close()

	entry, err := network.Lookup(ctx, host)
	io.FatalIfError(err)

	if entry != nil {
		fmt.Println(entry.Entry)
	} else {
		fmt.Println("host not found")
	}

}
