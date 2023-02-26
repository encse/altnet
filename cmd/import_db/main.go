package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/encse/altnet/ent"
	"github.com/encse/altnet/ent/schema"
	"github.com/encse/altnet/lib/io"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:./data/altnet.db?cache=shared&mode=rwc&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	client.Host.Delete().Exec(ctx)
	uumapBytes, err := ioutil.ReadFile("data/uumap.json")
	io.FatalIfError(err)

	var repr map[string]ent.Host
	err = json.Unmarshal(uumapBytes, &repr)
	io.FatalIfError(err)

	for _, v := range repr {
		h, err := client.Host.Create().
			SetName(v.Name).
			SetEntry(v.Entry).
			SetNeighbours(v.Neighbours).
			SetCountry(v.Country).
			SetMachineType(v.MachineType).
			SetOrganization(v.Organization).
			SetContact(v.Contact).
			SetContactAddress(v.ContactAddress).
			SetPhone(v.Phone).
			SetLocation(v.Location).
			SetGeoLocation(v.GeoLocation).
			Save(ctx)
		io.FatalIfError(err)
		fmt.Print(h.Name + "\r")
	}

	_, err = client.Host.Create().
		SetName(schema.HostName("csokavar")).
		SetNeighbours([]schema.HostName{
			"oddjob",
			"adaptex",
			"aemsrc",
			"bpsm",
			"tandem",
			"oracle",
			"veritas",
			"mimsy",
		}).
		Save(ctx)
	io.FatalIfError(err)
	fmt.Println("csokavar")
}
