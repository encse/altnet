package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/io"
)

func main() {
	ctx := altnet.ContextFromEnv(context.Background())
	host, err := altnet.GetHost(ctx)
	io.FatalIfError(err)

	fmt.Printf("HOST %s PROCESS TABLE\n", strings.ToUpper(string(host)))
	procinfos, err := altnet.GetProcesses(host)
	io.FatalIfError(err)
	fmt.Printf("%d running processes\n", len(procinfos))
	fmt.Println()

	table := [][]string{
		{"pid", "user", "started", "program"},
		{"---", "----", "-------", "-------"},
	}

	for _, procinfo := range procinfos {
		table = append(table,
			[]string{
				fmt.Sprintf("%v", procinfo.Pid),
				fmt.Sprintf("%v", procinfo.User),
				formatDate(procinfo.Started),
				fmt.Sprintf("%v", procinfo.Exe),
			},
		)
	}
	fmt.Println(io.Table(table...))
}

func formatDate(t time.Time) string {
	since := time.Now().Sub(t).Milliseconds()
	s := int64(1000)
	m := s * 60
	h := m * 60
	d := h * 24
	if since >= d {
		return fmt.Sprintf("%vd", since/d)
	}
	if since >= h {
		return fmt.Sprintf("%vh", since/h)
	}
	if since >= m {
		return fmt.Sprintf("%vm", since/m)
	}
	if since >= s {
		return fmt.Sprintf("%vs", since/s)
	}
	return fmt.Sprintf("0s")
}
