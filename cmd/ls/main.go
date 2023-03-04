package main

import (
	"context"
	"fmt"
	"syscall"

	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/slices"
	"golang.org/x/term"
)

func main() {
	ctx := altnet.ContextFromEnv(context.Background())

	files, err := altnet.Files(ctx)
	io.FatalIfError(err)

	columns, _, err := term.GetSize(int(syscall.Stdin))
	io.FatalIfError(err)

	names := slices.Map(files, func(fi altnet.FileInfo) string {
		return fi.Name()
	})

	maxWidth := 0
	for _, name := range names {
		if maxWidth < len(name) {
			maxWidth = len(name)
		}
	}

	lines := slices.Chunk(names, columns/(maxWidth+2))
	fmt.Print(io.Table(lines...))
}
