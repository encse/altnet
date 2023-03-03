package main

import (
	"context"
	"fmt"
	stdio "io"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/encse/altnet/ent"
	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/log"
	"github.com/encse/altnet/lib/slices"
	"github.com/encse/altnet/lib/uumap"
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

	fmt.Println(
		io.Center(
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
			), width))
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

	listFiles(files, width)

	for {
		item, err := io.Readline(fmt.Sprintf("Select file (%d-%d): ", 1, len(files)))
		idx := 0
		if err == nil {
			idx, err = strconv.Atoi(item)
		}

		if err != nil || idx < 1 || idx > len(files) {
			fmt.Println("ERROR: File not found.")
			return nil
		}

		for {
			option, err := io.Readline("Do you want to (P)rint, (D)ownload or (C)ancel? ")
			if err != nil {
				return err
			} else if strings.ToUpper(option) == "P" {
				printFile(ctx, files[idx].Name())
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
	return nil
}

func printFile(ctx context.Context, name string) error {
	fd, err := altnet.Open(ctx, name)
	io.FatalIfError(err)

	defer func() {
		err := fd.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	_, err = stdio.Copy(os.Stdout, fd)
	return err
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

	d := len(items) / 2
	for i := 0; i < d; i++ {
		fmt.Print(items[i])
		fmt.Print("  ")
		if i+d < len(items) {
			fmt.Print(items[i+d])
		}
		fmt.Println()
	}
}
