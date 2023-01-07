package main

import (
	"context"
	"fmt"

	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/io"
)

func main() {
	ctx := altnet.ContextFromEnv(context.Background())

	files, err := altnet.Files(ctx)
	io.FatalIfError(err)

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
