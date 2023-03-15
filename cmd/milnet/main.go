package main

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/encse/altnet/ent"
	"github.com/encse/altnet/ent/schema"
	"github.com/encse/altnet/ent/user"
	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/milnet"
	"github.com/encse/altnet/lib/slices"
	"github.com/encse/altnet/lib/uman"
	"github.com/encse/altnet/lib/uumap"
)

func main() {

	ctx := altnet.ContextFromEnv(context.Background())
	hostName, err := altnet.GetHost(ctx)
	io.FatalIfError(err)

	realUser, err := altnet.GetRealUser(ctx)
	io.FatalIfError(err)
	fmt.Println(realUser)
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
	fmt.Println("CONNECTED TO", strings.ToUpper(string(hostName)))
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
	valid := false

	if userid == realUser {
		c, err := host.QueryHackers().Where(user.UserEQ(realUser)).Count(ctx)
		io.FatalIfError(err)
		if c > 0 {
			valid, err = uman.ValidatePassword(ctx, network, realUser, schema.Password(password))
			io.FatalIfError(err)
		}
	} else {
		users, err := host.QueryVirtualusers().All(ctx)
		io.FatalIfError(err)

		valid = slices.Any(users, func(user *ent.VirtualUser) bool {
			return user.User == userid && user.Password == password
		})
	}

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
