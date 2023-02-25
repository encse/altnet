package main

import (
	"fmt"
	"sort"

	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/uumap"
)

func main() {
	network, err := uumap.GetUumap()
	io.FatalIfError(err)
	keys := network.Hosts()
	sort.Strings(keys)

	rows := make([][]string, 0, network.Size()+2)
	rows = append(rows, []string{"host", "organization", "location"})
	rows = append(rows, []string{"----", "------------", "--------"})

	for _, key := range keys {
		host, _ := network.Lookup(uumap.Host(key))
		rows = append(rows, []string{
			string(host.HostName),
			io.Substring(string(host.Organization), 32),
			io.Substring(string(host.Location), 32),
		})
	}

	fmt.Println(io.Table(rows...))
}
