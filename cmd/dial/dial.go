package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/phonenumbers"
	"github.com/encse/altnet/lib/uumap"
)

func main() {
	ctx := altnet.ContextFromEnv(context.Background())
	_, err := altnet.GetHost(ctx)
	io.FatalIfError(err)

	phonebook, err := uumap.GetPhonebook()
	io.FatalIfError(err)

	tel := ""
	if len(os.Args) > 1 {
		tel = strings.Join(os.Args[1:], " ")
	} else {
		tel, err = io.Readline("Tel:")
		io.FatalIfError(err)
	}

	phoneNumber, err := phonenumbers.ParsePhoneNumber(tel)
	if err != nil {
		fmt.Println("Invalid phone number.")
		return
	}

	atdt, err := phoneNumber.ToAtdtString()
	if err != nil {
		fmt.Println("Invalid phone number.")
		return
	}
	fmt.Print("  dialing ")
	io.SlowPrint(atdt)
	fmt.Print("     ")
	<-time.After(2 * time.Second)

	if host, ok := phonebook.Lookup(phoneNumber); ok {
		fmt.Println("CONNECT")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		altnet.Login(ctx, host)
		io.SlowPrint("?=\"[<}|}&'|!?+++ATH0\n")
		fmt.Println("NO CARRIER")
		fmt.Printf("%%disconnected\n")
	} else {
		fmt.Println("NO CARRIER")
	}
}
