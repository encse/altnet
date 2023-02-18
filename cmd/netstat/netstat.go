package main

import (
	"context"
	"fmt"
	"sort"

	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/slices"
	"github.com/encse/altnet/lib/uumap"
)

func main() {
	ctx := altnet.ContextFromEnv(context.Background())
	host, err := altnet.GetHost(ctx)
	io.FatalIfError(err)

	uumap, err := uumap.GetUumap()
	io.FatalIfError(err)

	uuentry, ok := uumap[string(host)]
	if !ok {
		fmt.Println("host not found")
		return
	}

	keys := slices.Clone(uuentry.Hosts)
	sort.Strings(keys)

	rows := make([][]string, 0, len(keys)+2)
	rows = append(rows, []string{"host", "organization", "location"})
	rows = append(rows, []string{"----", "------------", "--------"})

	for _, key := range keys {
		host := uumap[key]
		rows = append(rows, []string{
			string(host.HostName),
			io.Substring(string(host.Organization), 32),
			io.Substring(string(host.Location), 32),
		})
	}

	fmt.Println(io.Table(rows...))
}
