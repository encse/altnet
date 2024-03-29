package uumap

import (
	"context"

	"github.com/encse/altnet/ent"
	"github.com/encse/altnet/ent/host"
	"github.com/encse/altnet/schema"

	_ "github.com/mattn/go-sqlite3"
)

const connectionString = "file:./data/altnet.db?cache=shared&mode=rwc&_fk=1"

type Network struct {
	Client *ent.Client
}

func (n Network) Lookup(ctx context.Context, hostName schema.HostName) (*ent.Host, error) {

	host, err := n.Client.Host.
		Query().
		Where(host.Name(hostName)).
		First(ctx)
	return host, err
}

func (n Network) Close() error {
	return n.Client.Close()
}

func NetworkConn() (Network, error) {
	client, err := ent.Open("sqlite3", connectionString)
	return Network{Client: client}, err
}
