package main

import (
	"context"
	"fmt"
	"sort"

	"github.com/encse/altnet/ent/schema"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/uumap"
)

func main() {
	ctx := context.Background()
	network, err := uumap.NetworkConn()
	io.FatalIfError(err)
	defer network.Close()

	keys, err := network.Hosts(ctx)
	io.FatalIfError(err)
	sort.Strings(keys)

	rows := make([][]string, 0, len(keys)+2)
	rows = append(rows, []string{"host", "organization", "location"})
	rows = append(rows, []string{"----", "------------", "--------"})

	for _, key := range keys {
		host, _ := network.Lookup(ctx, schema.HostName(key))
		rows = append(rows, []string{
			string(host.Name),
			io.Substring(string(host.Organization), 32),
			io.Substring(string(host.Location), 32),
		})
	}

	fmt.Println(io.Table(rows...))
}
