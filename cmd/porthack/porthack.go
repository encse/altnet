package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/encse/altnet/ent"
	"github.com/encse/altnet/ent/user"
	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/slices"
	"github.com/encse/altnet/lib/uumap"
	"github.com/encse/altnet/schema"
)

func main() {

	fmt.Println("._   _  ._ _|_ |_   _.  _ |  ")
	fmt.Println("|_) (_) |   |_ | | (_| (_ |< ")
	fmt.Println("|                            ")
	fmt.Println("Starting Porthack 0.9 at", time.Now().Format(time.RFC1123))
	fmt.Println("")

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

	hostName, err := altnet.GetHost(ctx)
	io.FatalIfError(err)

	node, err := network.Lookup(ctx, hostName)
	io.FatalIfError(err)
	if node == nil {
		fmt.Println("host not found")
		return
	}

	targetHost := schema.HostName("")
	if len(os.Args) > 1 {
		targetHost = schema.HostName(os.Args[1])
	} else {
		targetHost, err = io.ReadArgFromList("host", os.Args, 1, node.Neighbours)
		io.FatalIfError(err)
	}

	if !slices.Contains(node.Neighbours, targetHost) {
		fmt.Println("host not in NETSTAT")
		return
	}

	targetNode, err := network.Lookup(ctx, targetHost)
	io.FatalIfError(err)

	fmt.Printf("Target: %s, initiating Connect Scan.\n", targetHost.ToUpper())

	c, err := network.Client.TcpService.Query().Count(ctx)
	io.FatalIfError(err)

	fmt.Printf("Scanning [%d] ports", c)
	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Println("completed.")
	fmt.Println("")

	err = exploit(ctx, network.Client, realUser, targetNode)

	io.FatalIfError(err)
}

func exploit(
	ctx context.Context,
	client *ent.Client,
	realUser schema.Uname,
	node *ent.Host,
) error {

	services, err := node.QueryServices().All(ctx)
	io.FatalIfError(err)

	rows := make([][]string, 0)
	rows = append(rows, []string{"port", "name", "description"})
	rows = append(rows, []string{"----", "----", "-----------"})

	for _, service := range services {
		rows = append(rows, []string{
			strconv.Itoa(service.Port),
			service.Name,
			service.Description,
		})
	}
	fmt.Println(io.Table(rows...))

	vulnerableService, err := slices.Choose(services)
	if err != nil {
		return err
	}

	for {
		rsp, err := io.Readline("Port to try: ")
		if err != nil {
			return err
		}

		found := false
		for _, service := range services {
			if rsp == strconv.Itoa(service.Port) {
				found = true
				fmt.Println("Scanning port", rsp)
				time.Sleep(2 * time.Second)
				if vulnerableService == service {
					fmt.Println("Segfault detected, building NOP sled...")
					time.Sleep(1 * time.Second)
					fmt.Println("Found libc base, preparing ROP chain.")
					time.Sleep(1 * time.Second)
					fmt.Println("Got shell, creating user.")
					time.Sleep(1 * time.Second)

					u, err := client.User.Query().Where(user.UserEQ(realUser)).First(ctx)
					io.FatalIfError(err)

					err = node.Update().AddHackers(u).Exec(ctx)
					io.FatalIfError(err)
					fmt.Printf(
						"Host %s hacked, you can TELNET to it and LOGIN with your regular credentials.\n",
						node.Name.ToUpper(),
					)
					return nil

				} else {
					fmt.Println("Port is not vulnerable")
				}
			}
		}
		if !found {
			fmt.Println("Invalid port.")
		}
	}
}
