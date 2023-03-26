package main

import (
	"context"
	"os"
	"regexp"

	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/bin"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/slices"
)

func main() {
	ctx := altnet.ContextFromEnv(context.Background())
	files, err := altnet.Files(ctx)
	io.FatalIfError(err)

	flags := slices.GetOrDefault(os.Args, 1)
	if ok, _ := regexp.Match("-.*l", []byte(flags)); ok {
		err = bin.LsWide(files)
	} else {
		err = bin.Ls(files)
	}
	io.FatalIfError(err)
}
