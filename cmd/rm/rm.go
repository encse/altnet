package main

import (
	"context"
	"os"

	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/io"
)

func main() {
	ctx := altnet.ContextFromEnv(context.Background())

	name, err := io.ReadArg("file", os.Args, 1)
	io.FatalIfError(err)

	err = altnet.Del(ctx, name)
	io.FatalIfError(err)
}
