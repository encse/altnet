package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/uumap"
)

func main() {
	hostname, err := io.ReadArg("host", os.Args, 1)
	io.FatalIfError(err)

	host := uumap.Host(strings.ToLower(hostname))

	network, err := uumap.GetUumap()
	io.FatalIfError(err)

	if entry, ok := network.Lookup(host); ok {
		fmt.Println(entry.Entry)
	} else {
		fmt.Println("host not found")
	}

}
