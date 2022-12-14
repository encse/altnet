package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/uumap"
)

func main() {
	host, err := io.ReadArg("host", os.Args, 1)
	io.FatalIfError(err)

	host = strings.ToLower(host)

	uumap, err := uumap.GetUumap()
	io.FatalIfError(err)

	if entry, ok := uumap[host]; ok {
		fmt.Println(entry.Entry)
	} else {
		fmt.Println("host not found")
	}

}
