package main

import (
	"context"
	"fmt"
	stdio "io"
	"os"

	iconv "github.com/djimenez/iconv-go"
	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/io"
	xtext "golang.org/x/text/transform"
)

func main() {

	ctx := altnet.ContextFromEnv(context.Background())

	name, err := io.ReadArg("file", os.Args, 1)
	io.FatalIfError(err)

	fi, err := altnet.GetFileInfo(ctx, name)
	io.FatalIfError(err)
	fmt.Print(io.ClearScreen + io.Home)

	in, err := os.Open(fi.RealPath())
	io.FatalIfError(err)

	f1 := xtext.NewReader(in, zeroTo32Transformer{})
	f2, err := iconv.NewReader(f1, "437", "UTF-8")
	io.FatalIfError(err)

	_, err = stdio.Copy(os.Stdout, f2)
	io.FatalIfError(err)

	io.ReadKey()
}

type zeroTo32Transformer struct{}

func (zeroTo32Transformer) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	for i := 0; i < len(src); i++ {
		if src[i] == 0 {
			src[i] = 32
		}
		dst[i] = src[i]
		nDst += 1
		nSrc += 1
	}
	return nDst, nSrc, nil
}

func (zeroTo32Transformer) Reset() {}
