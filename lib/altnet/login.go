package altnet

import (
	"context"
	"fmt"
	"strings"

	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/log"
)

func Login(ctx context.Context, host Host) {
	fmt.Println(fmt.Sprintf("Connected to %s", strings.ToUpper(string(host))))
	fmt.Println()
	fmt.Println("Enter your username or GUEST")

	username, err := io.ReadNotEmpty("Login: ")
	io.FatalIfError(err)

	username = strings.ToLower(username)
	if username != "guest" {
		for i := 0; i < 3; i++ {
			_, err = io.ReadPassword("Password: ")
			io.FatalIfError(err)
		}
	} else {
		ctx = SetHost(ctx, host)
		ctx = SetUser(ctx, User(username))

		log.Infof("Connected as %s", username)
		fmt.Println("Welcome", username)
		RunHiddenCommand(ctx, "./shell")
	}
}
