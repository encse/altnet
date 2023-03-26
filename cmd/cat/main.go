package main

import (
	"context"
	"os"

	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/bin"
	"github.com/encse/altnet/lib/io"
)

func main() {
	ctx := altnet.ContextFromEnv(context.Background())

	name, err := io.ReadArg("file", os.Args, 1)
	io.FatalIfError(err)

	fi, err := altnet.GetFileInfo(ctx, name)
	io.FatalIfError(err)
	bin.Cat(ctx, fi)
}
