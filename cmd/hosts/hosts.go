package main

import (
	"fmt"
	"sort"

	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/maps"
	"github.com/encse/altnet/lib/uumap"
)

func main() {
	uumap, err := uumap.GetUumap()
	io.FatalIfError(err)
	keys := maps.Keys(uumap)
	sort.Strings(keys)

	rows := make([][]string, 0, len(uumap)+2)
	rows = append(rows, []string{"host", "organization", "location"})
	rows = append(rows, []string{"----", "------------", "--------"})

	for _, key := range keys {
		host := uumap[key]
		rows = append(rows, []string{
			string(host.HostName),
			io.Substring(string(host.Organization), 36),
			io.Substring(string(host.Location), 18),
		})
	}

	fmt.Println(io.Table(rows...))
}
