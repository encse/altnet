package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/encse/altnet/ent"
	"github.com/encse/altnet/ent/host"
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

	_, err = client.Host.Delete().Exec(ctx)
	io.FatalIfError(err)

	err = importUumap(ctx, client)
	if err != nil {
		fmt.Println("could not import uumap hosts", err)
	}

	err = importBbs(ctx, client)
	if err != nil {
		fmt.Println("could not import bbs hosts", err)
	}

	err = importCsokavar(ctx, client)
	if err != nil {
		fmt.Println("could not import csokavar", err)
	}

	err = importJokes(ctx, client)
	if err != nil {
		fmt.Println("could not import jokes", err)
	}
}

func importJokes(ctx context.Context, client *ent.Client) error {
	jokesBytes, err := ioutil.ReadFile("data/jokes.json")
	io.FatalIfError(err)

	type entry struct {
		Id       int    `json:"id"`
		Body     string `json:"body"`
		Category string `json:"category"`
	}

	var entries []entry
	err = json.Unmarshal(jokesBytes, &entries)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if strings.TrimSpace(entry.Body) == "" {
			continue
		}

		_, err := client.Joke.Create().
			SetID(entry.Id).
			SetBody(entry.Body).
			SetCategory(entry.Category).
			Save(ctx)

		if err != nil {
			fmt.Println(entry.Id, err)
			continue
		}
	}
	return nil
}

func importBbs(ctx context.Context, client *ent.Client) error {
	bbsBytes, err := ioutil.ReadFile("data/bbs.json")
	io.FatalIfError(err)

	type entry struct {
		Phone    string `json:"phone"`
		Names    string `json:"names"`
		Location string `json:"location"`
		Sysop    string `json:"sysop"`
		YearFrom int    `json:"year_from"`
		YearTo   int    `json:"year_to"`
	}

	var entries []entry
	err = json.Unmarshal(bbsBytes, &entries)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		names := strings.Split(entry.Names, ",")
		if len(names) == 0 || names[0] == "" {
			continue
		}

		sysops := strings.Split(entry.Sysop, ",")
		if len(sysops) == 0 || sysops[0] == "" {
			continue
		}

		if entry.YearFrom > 1992 || entry.YearTo < 1992 {
			continue
		}

		if entry.Phone == "" {
			continue
		}

		name := schema.HostName(names[0])
		_, err := client.Host.Create().
			SetType(host.TypeBbs).
			SetName(name).
			SetContact(sysops[0]).
			SetPhone([]schema.PhoneNumber{schema.PhoneNumber(entry.Phone)}).
			SetLocation(entry.Location).
			Save(ctx)

		if err != nil {
			fmt.Println(name, err)
			continue
		}

		fmt.Print(names[0] + "                     \r")
	}
	return nil
}

func importUumap(ctx context.Context, client *ent.Client) error {

	uumapBytes, err := ioutil.ReadFile("data/uumap.json")
	if err != nil {
		return err
	}

	var repr map[string]ent.Host
	err = json.Unmarshal(uumapBytes, &repr)
	if err != nil {
		return err
	}

	for _, v := range repr {
		h, err := client.Host.Create().
			SetName(v.Name).
			SetType(host.TypeUucp).
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
		if err != nil {
			fmt.Println(v.Name, err)
			continue
		}
		fmt.Print(h.Name + "                     \r")
	}
	return nil
}

func importCsokavar(ctx context.Context, client *ent.Client) error {
	_, err := client.Host.Create().
		SetType(host.TypeUucp).
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
	return err
}
