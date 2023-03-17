package main

import (
	"context"
	"fmt"

	"github.com/encse/altnet/ent/user"
	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/uumap"
)

func main() {

	ctx := altnet.ContextFromEnv(context.Background())
	realUser, err := altnet.GetRealUser(ctx)
	io.FatalIfError(err)

	if realUser == "guest" {
		fmt.Println("Cannot run as guest, exiting.")
		return
	}

	network, err := uumap.NetworkConn()
	io.FatalIfError(err)
	defer network.Close()

	u, err := network.Client.User.Query().Where(user.UserEQ(realUser)).First(ctx)
	io.FatalIfError(err)

	hosts, err := u.QueryHosts().All(ctx)
	io.FatalIfError(err)
	if hosts != nil {
		fmt.Println("logins:")
		rows := make([][]string, 0, len(hosts)+2)
		rows = append(rows, []string{"host", "organization", "location"})
		rows = append(rows, []string{"----", "------------", "--------"})

		for _, host := range hosts {
			rows = append(rows, []string{
				string(host.Name),
				io.Linebreak(string(host.Organization), 32),
				io.Linebreak(string(host.Location), 32),
			})
		}
		fmt.Println(io.Table(rows...))
	} else {
		fmt.Printf("no logins")
	}
}
