package main

import (
	"context"
	"fmt"
	"strings"

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
		rows = append(rows, []string{"host", "phone", "organization"})
		rows = append(rows, []string{"----", "-----", "------------"})

		for _, host := range hosts {
			phone := ""
			if len(host.Phone) > 0 {
				phone = string(host.Phone[0])
			}

			org := strings.Split(host.Organization, "\n")[0]

			rows = append(rows, []string{
				string(host.Name),
				phone,
				io.Substring(org, 32),
			})
		}
		fmt.Println(io.Table(rows...))
	} else {
		fmt.Printf("no logins")
	}
}
