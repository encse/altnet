package main

import (
	"context"
	"fmt"

	"github.com/encse/altnet/ent/schema"
	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/slices"
	"github.com/encse/altnet/lib/uumap"
)

func main() {
	ctx := altnet.ContextFromEnv(context.Background())
	host, err := altnet.GetHost(ctx)
	io.FatalIfError(err)

	network, err := uumap.NetworkConn()
	io.FatalIfError(err)
	defer network.Close()

	uuentry, err := network.Lookup(ctx, host)
	io.FatalIfError(err)

	if uuentry == nil {
		fmt.Println("host not found")
		return
	}

	neighbours := slices.Clone(uuentry.Neighbours)
	slices.Sort(neighbours)

	rows := make([][]string, 0, len(neighbours)+2)
	rows = append(rows, []string{"host", "organization", "location"})
	rows = append(rows, []string{"----", "------------", "--------"})

	for _, key := range neighbours {
		host, _ := network.Lookup(ctx, schema.HostName(key))
		rows = append(rows, []string{
			string(host.Name),
			io.Substring(string(host.Organization), 32),
			io.Substring(string(host.Location), 32),
		})
	}

	fmt.Println(io.Table(rows...))
}
