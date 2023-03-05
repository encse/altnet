package altnet

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/encse/altnet/ent"
	"github.com/encse/altnet/ent/host"
	"github.com/encse/altnet/ent/schema"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/log"
	"github.com/encse/altnet/lib/uumap"
)

func Login(ctx context.Context, h *ent.Host) {

	ctx = SetHost(ctx, h.Name)

	if h.Type == host.TypeUucp {
		fmt.Printf("Connected to %s\n", strings.ToUpper(string(h.Name)))
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
			ctx = SetUser(ctx, User(username))
			log.Infof("Connected as %s", username)
			fmt.Println("Welcome", username)
			RunHiddenCommand(ctx, "./shell")
		}
	} else if h.Type == host.TypeBbs {
		RunHiddenCommand(ctx, "./datadrivebbs")
	} else {
		RunHiddenCommand(ctx, "./milnet")
	}
}

// Dial calls the given phone number in the phone book. If there is a host
// registered to that number, it tries to establish a connection with the host
// and starts a login session. The result is true. If there is host listening
// or the line is busy, dial returns false.
func Dial(
	ctx context.Context,
	phonenumber schema.PhoneNumber,
	network uumap.Network,
) (bool, error) {
	atdt, err := phonenumber.ToAtdtString()
	if err != nil {
		return false, err
	}

	fmt.Print("  dialing ")
	io.SlowPrint(atdt)
	fmt.Print("    ")
	time.Sleep(2 * time.Second)

	host, err := network.LookupHostByPhone(ctx, schema.PhoneNumber(phonenumber))
	if err != nil {
		return false, err
	}

	if host != nil {
		fmt.Println("CONNECT")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		Login(ctx, host)
		io.SlowPrint("?=\"[<}|}&'|!?+++ATH0\n")
		fmt.Println("NO CARRIER")
		fmt.Printf("%%disconnected\n")
		return true, nil
	} else {
		fmt.Println("NO CARRIER")
		return false, nil
	}
}
