package main

import (
	"context"
	"fmt"
	"time"

	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/uumap"
)

func main() {
	ctx := context.Background()
	network, err := uumap.NetworkConn()
	io.FatalIfError(err)
	defer network.Close()

	users, err := network.Client.User.Query().All(ctx)
	io.FatalIfError(err)
	rows := make([][]string, 0, len(users)+2)
	rows = append(rows, []string{"username", "last"})
	rows = append(rows, []string{"--------", "----"})

	for _, user := range users {
		lastLogin := "-"
		if user.LastLogin != nil {
			lastLogin = formatDate(*user.LastLogin)
		}
		rows = append(rows, []string{
			string(user.User),
			lastLogin,
		})
	}

	fmt.Println(io.Table(rows...))
}

func formatDate(t time.Time) string {
	since := time.Since(t).Milliseconds()
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
	return "0s"
}
