package main

import (
	"context"
	"fmt"

	"github.com/encse/altnet/ent"
	"github.com/encse/altnet/ent/host"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/uumap"
)

func main() {
	ctx := context.Background()
	network, err := uumap.NetworkConn()
	io.FatalIfError(err)
	defer network.Close()

	var hosts []struct {
		Name         string `json:"name,omitempty"`
		Organization string `json:"organization,omitempty"`
		Location     string `json:"location,omitempty"`
	}

	err = network.Client.Host.
		Query().
		Where(host.TypeIn(host.TypeUucp)).
		Order(ent.Asc(host.FieldName)).
		Select(host.FieldName, host.FieldOrganization, host.FieldLocation).
		Scan(ctx, &hosts)

	io.FatalIfError(err)

	rows := make([][]string, 0, len(hosts)+2)
	rows = append(rows, []string{"host", "organization", "location"})
	rows = append(rows, []string{"----", "------------", "--------"})

	for _, host := range hosts {
		rows = append(rows, []string{
			string(host.Name),
			io.Substring(string(host.Organization), 32),
			io.Substring(string(host.Location), 32),
		})
	}

	fmt.Println(io.Table(rows...))
}
