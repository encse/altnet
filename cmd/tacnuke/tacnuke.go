package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/encse/altnet/ent"
	"github.com/encse/altnet/ent/host"
	"github.com/encse/altnet/ent/user"
	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/milnet"
	"github.com/encse/altnet/lib/slices"
	"github.com/encse/altnet/lib/uumap"
	"github.com/encse/altnet/schema"
)

func main() {

	fmt.Println("")
	fmt.Println("         .-^^---....,,--        ")
	fmt.Println("     _--                  --_   ")
	fmt.Println("    <       T.A.C NUKE       >  ")
	fmt.Println("    |       (C) encse         | ")
	fmt.Println("     \\._                   _./  ")
	fmt.Println("        ```--. . , ; .--'''      ")
	fmt.Println("              | |   |            ")
	fmt.Println("           .-=||  | |=-.         ")
	fmt.Println("           `-=#$%&%$#=-'         ")
	fmt.Println("              | ;  :|            ")
	fmt.Println("     _____.,-#%&$@%#&#~,._____   ")
	fmt.Println("")
	io.Readline("Press enter to continue")
	fmt.Println("")

	ctx := altnet.ContextFromEnv(context.Background())

	userName, err := altnet.GetUser(ctx)
	io.FatalIfError(err)

	if userName == "guest" {
		fmt.Println("Cannot run as guest, exiting.")
		return
	}

	network, err := uumap.NetworkConn()
	io.FatalIfError(err)
	defer network.Close()

	tel := ""
	if len(os.Args) > 1 {
		tel = strings.Join(os.Args[1:], " ")
	} else {
		tel, err = io.Readline("Tel:")
		io.FatalIfError(err)
	}

	phoneNumber, err := schema.ParsePhoneNumber(tel)
	if err != nil {
		fmt.Println("Invalid phone number.")
		return
	}

	fmt.Print("Dialing " + phoneNumber.ToUsLocalString() + "    ")

	node, err := network.LookupHostByPhone(ctx, phoneNumber)
	io.FatalIfError(err)

	if node == nil {
		fmt.Println("NO CARRIER")
		return
	} else {
		fmt.Println("CONNECTED")
	}

	users, err := node.QueryVirtualusers().All(ctx)
	io.FatalIfError(err)

	fmt.Println("Looking for TAC header...")
	time.Sleep(500 * time.Millisecond)
	if node.Type != host.TypeMil {
		fmt.Println("No header found, not a TAC login prompt.")
		return
	}

	fmt.Print("Attempting to exploit TAC CRC32")
	for i := 0; i < 6; i++ {
		fmt.Print(".")
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("   successful")

	fmt.Println("Dumping memory")

	generateDump(ctx, node)
	for {

		fmt.Println("")

		userIdEntered, err := io.Readline("Enter user id: ")
		io.FatalIfError(err)

		if userIdEntered == "r" || userIdEntered == "retry" {
			generateDump(ctx, node)
			continue
		}
		if userIdEntered == "q" || userIdEntered == "quit" {
			break
		}

		passwordEntered, err := io.Readline("Enter password: ")
		io.FatalIfError(err)

		fmt.Print("Validating credentials...")
		time.Sleep(800 * time.Millisecond)

		valid := slices.Any(users, func(user *ent.VirtualUser) bool {
			return user.User == schema.Uname(userIdEntered) && user.Password == schema.Password(passwordEntered)
		})

		if valid {
			fmt.Println("ok.")
			fmt.Printf("Sending user registration request in the name of '%s'... ok.\n", userIdEntered)
			time.Sleep(1000 * time.Millisecond)

			realUser, err := altnet.GetRealUser(ctx)
			io.FatalIfError(err)

			u, err := network.Client.User.Query().Where(user.UserEQ(realUser)).First(ctx)
			io.FatalIfError(err)

			err = node.Update().AddHackers(u).Exec(ctx)
			io.FatalIfError(err)

			fmt.Println("Access granted, use your regular login credentials to connect to the host.")
			break
		} else {
			fmt.Println("failed.")
		}
	}

	fmt.Println("CONNECTION CLOSED")
}

func generateDump(ctx context.Context, node *ent.Host) {

	mem := generateMemoryDump(ctx, node)

	st := ""
	for i := 0; i < len(mem); i++ {
		if (i % 16) == 0 {
			fmt.Printf("%04x  ", i)
		}

		fmt.Printf("%02x ", mem[i])
		if mem[i] >= 32 && mem[i] < 127 {
			st += fmt.Sprintf("%c", mem[i])
		} else {
			st += "."
		}

		if (i % 16) == 15 {
			fmt.Println("   " + st)
			st = ""
		}
	}

	fmt.Println("")
	fmt.Println("Use the TAC memory dump to extract a user-id and an access-code for the next step.")
	fmt.Println("Enter 'retry' or 'r' to generate a new dump, 'quit' or 'q' to exit")
}

func generateMemoryDump(ctx context.Context, host *ent.Host) []byte {

	mem := generateGarbage()
	mem = addStrings(mem, host.Location)
	mem = addStrings(mem, string(host.Name))

	for i := 0; i < 13; i++ {
		mem = addStrings(mem, milnet.RandomWord())
	}

	if rand.Float32() < 0.2 {
		virtualusers, _ := host.QueryVirtualusers().All(ctx)
		if len(virtualusers) > 0 {
			user := slices.ChooseX(virtualusers)
			mem = addStrings(
				mem,
				milnet.UserIdPrompt,
				milnet.PasswordPrompt,
				string(user.User),
				string(user.Password),
			)
		}
	}

	for i := 0; i < 13; i++ {
		mem = addStrings(mem, milnet.RandomWord())
	}

	return mem

}

func addStrings(mem []byte, args ...string) []byte {
	data := make([]byte, 0)
	for _, arg := range args {
		data = append(data, byte(len(arg)))
		data = append(data, []byte(arg)...)
	}

	ptr := rand.Intn(len(mem) - len(data))

	for _, b := range data {
		mem[ptr] = b
		ptr++
	}
	return mem
}

func generateGarbage() []byte {

	size := 16 * (30 + rand.Intn(10))
	var bytes = make([]byte, size)
	for i := 0; i < size; i++ {
		bytes[i] = byte(rand.Intn(256))
	}

	return bytes
}
