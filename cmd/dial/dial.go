package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/encse/altnet/ent/schema"
	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/uumap"
)

func main() {
	ctx := altnet.ContextFromEnv(context.Background())
	_, err := altnet.GetHost(ctx)
	io.FatalIfError(err)

	network, err := uumap.NetworkConn()
	io.FatalIfError(err)
	defer network.Close()

	tel := ""
	if len(os.Args) > 1 {
		tel = strings.Join(os.Args[1:], " ")
	} else {
		tel, err = io.Readline("Tel:")
		io.FatalIfError(err)
	}

	phoneNumber, err := schema.ParsePhoneNumber(tel)
	if err != nil {
		fmt.Println("Invalid phone number.")
		return
	}

	_, err = altnet.Dial(ctx, phoneNumber, network)
	io.FatalIfError(err)
}
