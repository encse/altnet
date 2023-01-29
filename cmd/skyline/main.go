package main

import (
	"fmt"
	"os"

	"github.com/encse/altnet/lib/csokavar"
	"github.com/encse/altnet/lib/io"
)

func main() {
	user, err := io.ReadArg("user", os.Args, 1)
	io.FatalIfError(err)

	st, err := csokavar.GetSkyline(user, 80)
	io.FatalIfError(err, fmt.Sprintf("Cannnot get skyline for %s now.", user))

	fmt.Println(st)
}
