package altnet

import (
	"context"
	"fmt"
	"time"

	"github.com/encse/altnet/ent"
	"github.com/encse/altnet/ent/host"
	"github.com/encse/altnet/ent/user"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/slices"
	"github.com/encse/altnet/lib/uman"
	"github.com/encse/altnet/lib/uumap"
	"github.com/encse/altnet/schema"
)

func ValidatePassword(
	ctx context.Context,
	network uumap.Network,
	h *ent.Host,
	userid schema.Uname,
	password schema.Password,
) (bool, error) {
	realUser, err := GetRealUser(ctx)
	io.FatalIfError(err)
	if userid == realUser {
		c, err := h.QueryHackers().Where(user.UserEQ(realUser)).Count(ctx)
		io.FatalIfError(err)
		if c > 0 {
			return uman.ValidatePassword(ctx, network, realUser, schema.Password(password))
		} else {
			return false, nil
		}
	} else {
		users, err := h.QueryVirtualusers().All(ctx)
		io.FatalIfError(err)

		valid := slices.Any(users, func(user *ent.VirtualUser) bool {
			return user.User == userid && user.Password == password
		})
		return valid, nil
	}
}

func Login(ctx context.Context, h *ent.Host) {
	ctx = SetHost(ctx, h.Name)
	if h.Type == host.TypeUucp {
		RunHiddenCommand(ctx, "./uucplogin")
	} else if h.Type == host.TypeBbs {
		RunHiddenCommand(ctx, "./datadrivebbs")
	} else if h.Type == host.TypeMil {
		RunHiddenCommand(ctx, "./milnetlogin")
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
