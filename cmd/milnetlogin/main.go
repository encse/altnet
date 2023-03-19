package main

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/milnet"
	"github.com/encse/altnet/lib/uumap"
	"github.com/encse/altnet/schema"
)

func main() {

	ctx := altnet.ContextFromEnv(context.Background())
	hostName, err := altnet.GetHost(ctx)
	io.FatalIfError(err)

	network, err := uumap.NetworkConn()
	io.FatalIfError(err)
	defer network.Close()

	host, err := network.Lookup(ctx, hostName)
	io.FatalIfError(err)

	for i := 0; i < 100; i++ {
		fmt.Print(strings.Repeat(" ", rand.Intn(800)))
		time.Sleep(30 * time.Millisecond)
		fmt.Print(milnet.RandomWord())
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("CONNECTED TO", hostName.ToUpper())
	fmt.Println(host.Location)
	fmt.Println()

	fmt.Println("WELCOME TO DDN. FOR OFFICIAL USE ONLY. TAC LOGIN REQUIRED.")
	fmt.Println("CALL NIC 1-800-235-3155 FOR HELP")
	fmt.Println()

	userid, password, err := readLoginCredentials(ctx)
	io.FatalIfError(err)

	fmt.Println()
	fmt.Print("Contacting authentication service...")
	time.Sleep(1 * time.Second)

	valid, err := altnet.ValidatePassword(ctx, network, host, userid, password)
	io.FatalIfError(err)

	if valid {
		fmt.Println("ok")
		fmt.Println()
		altnet.RunHiddenCommand(ctx, "./shell")
	} else {
		fmt.Println("denied")
		fmt.Println()
		fmt.Println("Login FAILED, closing connection.")
	}
}

func readLoginCredentials(ctx context.Context) (schema.Uname, schema.Password, error) {
	user, err := io.Readline(milnet.UserIdPrompt)
	if err != nil {
		return "", "", err
	}
	password, err := io.ReadPassword(milnet.PasswordPrompt)
	if err != nil {
		return "", "", err
	}
	return schema.Uname(user), schema.Password(password), nil
}
