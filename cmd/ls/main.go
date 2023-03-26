package main

import (
	"context"

	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/bin"
	"github.com/encse/altnet/lib/io"
)

func main() {
	ctx := altnet.ContextFromEnv(context.Background())
	files, err := altnet.Files(ctx)
	err = bin.Ls(files)
	io.FatalIfError(err)
}
