package main

import (
	"context"
	"os"

	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/csokavar"
	"github.com/encse/altnet/lib/io"
)

func main() {
	ctx := altnet.ContextFromEnv(context.Background())

	name, err := io.ReadArg("file", os.Args, 1)
	io.FatalIfError(err)

	fi, err := altnet.GetFileInfo(ctx, name)
	io.FatalIfError(err)
	csokavar.RunHiddenCommand(ctx, "/usr/bin/iconv", "-f", "437", "-t", "UTF-8", fi.RealPath())
}
