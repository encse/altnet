package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/uumap"
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

	fmt.Println(
		io.Box(
			"",
			string(host.Name),
			"",
			"10 gigs    4 lines",
			"",
			"Call "+string(host.Phone[0]),
			"",
			"SupraFAXModem V.32bis, DataDrive BBS 3600",
			"",
			"SYSOP: "+host.Contact,
			"",
			host.Location,
			"",
		))
	fmt.Println()

	fmt.Println("Enter your username or GUEST")
	username, err := io.ReadNotEmpty("Username: ")
	io.FatalIfError(err)

	username = strings.ToLower(username)
	if username != "guest" {
		for i := 0; i < 3; i++ {
			_, err = io.ReadPassword("Password: ")
			io.FatalIfError(err)
		}
		return
	}

	ctx = altnet.SetUser(ctx, altnet.User(username))
loop:

	for {
		fmt.Println()
		fmt.Println("<F> Files area        <W> Who's online")
		fmt.Println("<J> Tell a joke       <Y> Yell for sysop")
		fmt.Println("<Q> Quit / Logoff")
		fmt.Println("<?> Help")
		option, err := io.ReadOption("Select an item", "fwjyq?")
		io.FatalIfError(err)

		switch strings.ToLower(option) {
		case "f":
			fmt.Println("files")
		case "g":
			fmt.Println("who")
		case "j":
			joke, err := network.Joke(ctx)
			io.FatalIfError(err)
			fmt.Println()
			fmt.Println(io.Linebreak(joke, 80))
		case "y":
			fmt.Print("Pagig sysop...")
			time.Sleep(300 * time.Millisecond)
			fmt.Print(" ^G")
			time.Sleep(300 * time.Millisecond)
			fmt.Print(" ^G")
			time.Sleep(300 * time.Millisecond)
			fmt.Print(" ^G")
			time.Sleep(10 * time.Second)
			fmt.Println()
			fmt.Println()
			fmt.Println("No answer from sysop. Please try again later.")
		case "q":
			break loop
		}
	}
}
