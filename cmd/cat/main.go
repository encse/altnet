package main

import (
	"context"
	stdio "io"
	"os"

	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/log"
)

func main() {
	ctx := altnet.ContextFromEnv(context.Background())

	name, err := io.ReadArg("file", os.Args, 1)
	io.FatalIfError(err)

	fd, err := altnet.Open(ctx, name)
	io.FatalIfError(err)

	defer func() {
		err := fd.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	_, err = stdio.Copy(os.Stdout, fd)
	io.FatalIfError(err)
}
