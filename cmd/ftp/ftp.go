package main

import (
	"context"
	"errors"
	"fmt"
	stdio "io"
	"os"
	"strings"
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

	ctx := altnet.ContextFromEnv(context.Background())
	hostName, err := altnet.GetHost(ctx)
	io.FatalIfError(err)

	localuser, err := altnet.GetUser(ctx)
	io.FatalIfError(err)

	network, err := uumap.NetworkConn()
	io.FatalIfError(err)
	defer network.Close()

	localhost, err := network.Lookup(ctx, hostName)
	io.FatalIfError(err)

	mainLoop(ctx, network, localhost, localuser, slices.GetOrDefault(os.Args, 1))
}

type Ftp struct {
	localhost *ent.Host
	lctx      context.Context
	rctx      context.Context
	mode      string
}

func mainLoop(
	ctx context.Context,
	network uumap.Network,
	localhost *ent.Host,
	localuser schema.Uname,
	connectToHost string,
) error {
	ftp := Ftp{
		localhost: localhost,
		lctx:      ctx,
		rctx:      altnet.SetHost(altnet.SetUser(ctx, ""), ""),
	}

	if connectToHost != "" {
		var err error
		ftp, err = ftp.open(ctx, network, connectToHost)
		if err != nil {
			return err
		}
	}

	screen := struct {
		stdio.Reader
		stdio.Writer
	}{os.Stdin, os.Stdout}

	t := term.NewTerminal(screen, "ftp> ")
	for {
		cmd, err := readline(t)
		if err != nil {
			return err
		}

		parts := strings.Fields(cmd)
		if len(parts) == 0 {
			continue
		}

		switch strings.ToLower(parts[0]) {
		case "quit":
			return nil
		case "ascii":
			ftp, err = ftp.ascii(ctx, network)
		case "bin":
			ftp, err = ftp.bin(ctx, network)
		case "open":
			ftp, err = ftp.open(ctx, network, slices.GetOrDefault(parts, 1))
		case "del":
			ftp, err = ftp.del(ctx, network, slices.GetOrDefault(parts, 1))
		case "put":
			ftp, err = ftp.put(ctx, network, slices.GetOrDefault(parts, 1))
		case "get":
			ftp, err = ftp.get(ctx, network, slices.GetOrDefault(parts, 1))
		case "close", "bye":
			ftp, err = ftp.close(ctx, network)
		case "user":
			ftp, err = ftp.login(ctx, network, slices.GetOrDefault(parts, 1))
		case "ls":
			ftp, err = ftp.ls(ctx, network)
		case "lls":
			ftp, err = ftp.lls(ctx, network)
		case "help":
			fmt.Println("ascii   switch to ASCII mode")
			fmt.Println("bin     switch to binary mode")
			fmt.Println("close   close the connection")
			fmt.Println("del     delete files from remote host")
			fmt.Println("get     get files from remote host")
			fmt.Println("help    show ftp commands")
			fmt.Println("lls     list files on local host")
			fmt.Println("ls      list files on remote host")
			fmt.Println("open    connect to a remote host")
			fmt.Println("put     put files to remote host")
			fmt.Println("quit    exit ftp")
			fmt.Println("user    authenticate to the remote server")
		default:
			fmt.Println("502 Command not implemented")
		}

		if err != nil {
			return err
		}
	}
}

func readline(t *term.Terminal) (string, error) {
	oldState, err := term.MakeRaw(0)
	if err != nil {
		return "", err
	}

	defer func() {
		err := term.Restore(0, oldState)
		io.FatalIfError(err)
	}()
	return t.ReadLine()
}

func (ftp Ftp) ascii(ctx context.Context, network uumap.Network) (Ftp, error) {
	if !ftp.requireRemoteHost() {
		return ftp, nil
	}

	fmt.Println("200 Type set to A")
	ftp.mode = "ASCII"
	return ftp, nil
}

func (ftp Ftp) bin(ctx context.Context, network uumap.Network) (Ftp, error) {

	if !ftp.requireRemoteHost() {
		return ftp, nil
	}

	fmt.Println("200 Type set to I")
	ftp.mode = "BIN"
	return ftp, nil
}

func (ftp Ftp) close(ctx context.Context, network uumap.Network) (Ftp, error) {
	remotehost, err := altnet.GetHost(ftp.rctx)
	if err != nil {
		return ftp, err
	}

	if remotehost != "" {
		fmt.Println("221 Goodbye.")
		ftp.rctx = altnet.SetHost(altnet.SetUser(ctx, ""), "")
	}
	return ftp, nil
}

func (ftp Ftp) open(ctx context.Context, network uumap.Network, stHostName string) (Ftp, error) {
	var err error

	if stHostName == "" {
		stHostName, err = io.Readline("(to) ")
		if err != nil {
			return ftp, err
		}
	}

	stHostName = strings.TrimSpace(stHostName)
	if stHostName == "" {
		return ftp, nil
	}

	remoteHostName := schema.HostName(stHostName)
	if !slices.Contains(ftp.localhost.Neighbours, remoteHostName) {
		fmt.Printf("ftp: %s: host not in NETSTAT\n", remoteHostName)
		return ftp, nil
	}

	ftp.rctx = altnet.SetHost(altnet.SetUser(ctx, ""), remoteHostName)
	ftp.mode = "BIN"

	time.Sleep(1 * time.Second)

	fmt.Printf("Connected to %s.\n", remoteHostName.ToUpper())
	fmt.Printf("220 %s FTP server (Version 4.109 Wed Nov 19 21:52:18 CST 1986) ready.\n", remoteHostName)

	ftp, err = ftp.login(ctx, network, "")

	fmt.Println("Remote system type is UNIX.")
	fmt.Printf("Using %s mode to transfer files.\n", ftp.mode)
	return ftp, err
}

func (ftp Ftp) login(ctx context.Context, network uumap.Network, remoteuser string) (Ftp, error) {
	if !ftp.requireRemoteHost() {
		return ftp, nil
	}

	localuser, err := altnet.GetUser(ftp.lctx)
	if err != nil {
		return ftp, err
	}

	remotehost, err := ftp.getRemoteHost()
	if err != nil {
		return ftp, err
	}

	if remoteuser == "" {
		remoteuser, err = io.Readline(
			fmt.Sprintf("Name (%s:%s):", remotehost, localuser),
		)
		if err != nil {
			return ftp, err
		}
		remoteuser = strings.TrimSpace(remoteuser)
	}

	if remoteuser == "" {
		remoteuser = string(localuser)
	}

	valid := remoteuser == "guest"
	if !valid {
		fmt.Printf("331 Password required for %s\n", remoteuser)
		password, err := io.ReadPassword("Password: ")
		if err != nil {
			return ftp, err
		}

		h, err := network.Lookup(ctx, remotehost)
		if err != nil {
			return ftp, err
		}
		valid, err = altnet.ValidatePassword(
			ctx, network, h, schema.Uname(remoteuser), password,
		)

		if err != nil {
			return ftp, err
		}
	}

	if !valid {
		fmt.Println("530 Invalid username or password.")
		fmt.Println("Login failed.")
	} else {
		fmt.Printf("230 User %s logged in\n", remoteuser)
		ftp.rctx = altnet.SetUser(ftp.rctx, schema.Uname(remoteuser))
	}

	return ftp, nil
}

func (ftp Ftp) lls(ctx context.Context, network uumap.Network) (Ftp, error) {
	files, err := altnet.Files(ftp.lctx)
	if err != nil {
		return ftp, err
	}

	return ftp, bin.LsWide(files)
}

func (ftp Ftp) ls(ctx context.Context, network uumap.Network) (Ftp, error) {
	if !ftp.requireRemoteLogin() {
		return ftp, nil
	}

	files, err := altnet.Files(ftp.rctx)
	if err != nil {
		return ftp, err
	}

	fmt.Println("150 Opening ASCII mode data connection for file list")
	defer fmt.Println("226 Transfer complete")
	return ftp, bin.LsWide(files)
}

func (ftp Ftp) del(ctx context.Context, network uumap.Network, fileName string) (Ftp, error) {
	if !ftp.requireRemoteLogin() {
		return ftp, nil
	}

	fmt.Println("200 PORT command successful")
	err := altnet.Del(ftp.rctx, fileName)

	if uerr, ok := err.(io.UserFriendlyError); ok {
		fmt.Printf("550 %s: %s\n", fileName, &uerr)
		return ftp, nil
	} else if err != nil {
		return ftp, err
	} else {
		fmt.Println("250 DELE command successful")
		return ftp, nil
	}
}

func (ftp Ftp) put(ctx context.Context, network uumap.Network, fileName string) (Ftp, error) {
	return ftp.transfer(ftp.lctx, ftp.rctx, fileName, "sent")
}

func (ftp Ftp) get(ctx context.Context, network uumap.Network, fileName string) (Ftp, error) {
	return ftp.transfer(ftp.rctx, ftp.lctx, fileName, "received")
}

func (ftp Ftp) transfer(
	src context.Context,
	dst context.Context,
	fileName string,
	dir string,

) (Ftp, error) {
	if !ftp.requireRemoteLogin() {
		return ftp, nil
	}

	fmt.Println("200 PORT command successful")
	fi, err := altnet.GetFileInfo(src, fileName)
	if uerr, ok := err.(io.UserFriendlyError); ok {
		fmt.Printf("550 %s: %s\n", fileName, &uerr)
		return ftp, nil
	}
	fmt.Printf("150 Opening BIN mode data connection for %s (%v bytes)\n", fileName, fi.Size())

	now := time.Now()
	fakeProgress(30, time.Duration(fi.Size()/4096)*time.Second)

	err = altnet.Copy(src, dst, fileName)
	if uerr, ok := err.(io.UserFriendlyError); ok {
		fmt.Printf("550 %s: %s\n", fileName, &uerr)
		return ftp, nil
	} else if err != nil {
		return ftp, err
	}

	fmt.Println("226 Transfer complete")
	dt := time.Since(now).Seconds()

	seconds := int(dt)
	speed := float64(fi.Size()) / float64(dt) / 1000
	fmt.Printf("%v bytes %s in %d secs (%.02f kB/s)\n", fi.Size(), dir, seconds, speed)
	return ftp, nil
}

func fakeProgress(width int, duration time.Duration) {
	print := func(pct int) {
		done := strings.Repeat("#", int(width*pct/100.0))
		pad := strings.Repeat(".", width-len(done))
		fmt.Printf("\r%3d%% |%s%s|", pct, done, pad)
	}

	for pct := 0.0; pct < 100; pct += 10 / duration.Seconds() {
		print(int(pct))
		time.Sleep(100 * time.Millisecond)
	}
	print(100)
	fmt.Println()
}

func (ftp Ftp) getRemoteHost() (schema.HostName, error) {
	remotehost, err := altnet.GetHost(ftp.rctx)
	if err != nil {
		return "", err
	}
	if remotehost == "" {
		return "", errors.New("not connected")
	}

	return remotehost, nil
}

func (ftp Ftp) requireRemoteHost() bool {
	_, err := ftp.getRemoteHost()
	if err != nil {
		fmt.Println("Not connected.")
		return false
	}
	return true
}

func (ftp Ftp) requireRemoteLogin() bool {
	if !ftp.requireRemoteHost() {
		return false
	}
	remoteuser, err := altnet.GetUser(ftp.rctx)
	if err != nil {
		return false
	}

	if remoteuser == "" {
		fmt.Println("Need to LOGIN first.")
		return false
	}

	return true
}
