package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/encse/altnet/ent"
	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/bin"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/slices"
	"github.com/encse/altnet/lib/uumap"
	"github.com/encse/altnet/schema"
	"golang.org/x/term"
)

func main() {

	width, _, err := term.GetSize(int(syscall.Stdin))
	if err != nil {
		io.FatalIfError(err)
	}

	if width > 80 {
		width = 80
	}

	ctx := altnet.ContextFromEnv(context.Background())
	hostName, err := altnet.GetHost(ctx)
	io.FatalIfError(err)

	network, err := uumap.NetworkConn()
	io.FatalIfError(err)
	defer network.Close()

	host, err := network.Lookup(ctx, hostName)
	io.FatalIfError(err)

	fmt.Printf("Connected to %s.\n", host.Name.ToUpper())
	fmt.Println()
	fmt.Println(
		io.Center(
			io.Box(
				"",
				string(host.Organization),
				host.Phone[0].ToUsLocalString(),
				"",
				"SupraFAXModem V.32bis, DataDrive BBS 3600",
				"10 gigs    4 lines",
				"",
				"SYSOP: "+host.Contact,
				host.Location,
				"",
			), width))
	fmt.Println()

	fmt.Println("Enter your username or GUEST")
	username, err := io.ReadNotEmpty[schema.Uname]("Username: ")
	io.FatalIfError(err)

	username = username.ToLower()
	valid := username == "guest"
	for i := 0; i < 3 && !valid; i++ {
		password, err := io.ReadPassword("Password: ")
		valid, err = altnet.ValidatePassword(
			ctx, network, host, username, password,
		)
		io.FatalIfError(err)
	}

	if !valid {
		return
	}

	ctx = altnet.SetUser(ctx, username)

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
			err = filesArea(ctx, host, width)
			io.FatalIfError(err)
		case "g":
			fmt.Println("who")
		case "j":
			joke, err := network.Joke(ctx)
			io.FatalIfError(err)
			fmt.Println()
			fmt.Println(io.Linebreak(joke, 80))
		case "y":
			fmt.Print("Paging sysop...")
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

func filesArea(ctx context.Context, host *ent.Host, width int) error {
	files, err := altnet.Files(ctx)
	if err != nil {
		return err
	}

	for {
		listFiles(files, width)

		item, err := io.Readline(fmt.Sprintf("Select file (%d-%d): ", 1, len(files)))
		idx := 0
		if err == nil {
			idx, err = strconv.Atoi(item)
		}

		if err != nil || idx < 1 || idx > len(files) {
			fmt.Println("ERROR: File not found.")
			return nil
		}

		idx--

		for {
			option, err := io.Readline("Do you want to (P)rint, (D)ownload or (C)ancel? ")
			if err != nil {
				return err
			} else if strings.ToUpper(option) == "P" {
				printFile(ctx, files[idx].Name())
				fmt.Println()
				break
			} else if strings.ToUpper(option) == "D" {
				downloadFile(ctx, files[idx].Name())
				break
			} else if strings.ToUpper(option) == "C" {
				return nil
			}
		}
	}
}

func downloadFile(ctx context.Context, name string) error {
	fmt.Println("XMODEM transfer is ready to begin.")
	fmt.Println("Connecting...")
	time.Sleep(2 * time.Second)
	fmt.Println("Failed to connect, check that XMODEM is running on your host.")
	fmt.Println("Press <any> key to continue.")
	io.ReadKey()
	return nil
}

func printFile(ctx context.Context, name string) error {
	fi, err := altnet.GetFileInfo(ctx, name)
	if err != nil {
		return err
	}

	bin.More(ctx, fi)
	return nil
}

func listFiles(files []altnet.FileInfo, width int) {

	maxIndexWidth := len(strconv.Itoa(len(files)))
	maxSizeWidth := len(strconv.FormatInt(
		slices.Max(
			slices.Map(
				files,
				func(file altnet.FileInfo) int64 {
					return file.Size()
				},
			)),
		10))

	width = width/2 - 1

	maxNameWidth := width - maxIndexWidth - maxSizeWidth - 9

	items := make([]string, 0, len(files))
	for i, file := range files {
		row := fmt.Sprintf("%*d. %-*s %*d bytes",
			maxIndexWidth, i+1,
			maxNameWidth, strings.ToUpper(file.Name()),
			maxSizeWidth, file.Size(),
		)
		items = append(items, row)
	}

	d := (len(items) + 1) / 2
	for i := 0; i < d; i++ {
		fmt.Print(items[i])
		fmt.Print("  ")
		if i+d < len(items) {
			fmt.Print(items[i+d])
		}
		fmt.Println()
	}
}
