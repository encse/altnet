package uumap

import (
	"context"

	"github.com/encse/altnet/ent"
	"github.com/encse/altnet/ent/host"
	"github.com/encse/altnet/ent/schema"

	_ "github.com/mattn/go-sqlite3"
)

const connectionString = "file:./data/altnet.db?cache=shared&mode=rwc&_fk=1"

type Network struct {
	client *ent.Client
}

func (n Network) Hosts(ctx context.Context) ([]string, error) {
	var vs []struct {
		Name string `json:"name,omitempty"`
	}

	err := n.client.Host.
		Query().
		Select(host.FieldName).
		Scan(ctx, &vs)

	if err != nil {
		return nil, err
	}
	hosts := make([]string, 0, len(vs))
	for _, v := range vs {
		hosts = append(hosts, v.Name)
	}
	return hosts, nil
}

func (n Network) Lookup(ctx context.Context, hostName schema.HostName) (*ent.Host, error) {

	hosts, err := n.client.Host.
		Query().
		Where(host.Name(hostName)).
		All(ctx)

	if err != nil {
		return nil, err
	}

	if len(hosts) == 0 {
		return nil, nil
	}

	return hosts[0], nil
}

func (n Network) Close() error {
	return n.client.Close()
}

func NetworkConn() (Network, error) {
	client, err := ent.Open("sqlite3", connectionString)
	return Network{client: client}, err
}
