package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/encse/altnet/ent"
	"github.com/encse/altnet/ent/host"
	"github.com/encse/altnet/ent/tcpservice"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/milnet"
	"github.com/encse/altnet/lib/slices"
	"github.com/encse/altnet/schema"
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

	err = importServices(ctx, client)
	if err != nil {
		fmt.Println("could not import services", err)
	}

	err = importUumap(ctx, client)
	if err != nil {
		fmt.Println("could not import uumap hosts", err)
	}

	err = importBbs(ctx, client)
	if err != nil {
		fmt.Println("could not import bbs hosts", err)
	}

	err = importMilHosts(ctx, client)
	if err != nil {
		fmt.Println("could not import mil hosts", err)
	}

	err = importCsokavar(ctx, client)
	if err != nil {
		fmt.Println("could not import csokavar", err)
	}

	err = importJokes(ctx, client)
	if err != nil {
		fmt.Println("could not import jokes", err)
	}

	// Steve Jacksons's BBS
	client.Host.
		Update().
		Where(host.Name(schema.HostName("fnordbox"))).
		SetType(host.TypeBbs).
		Exec(ctx)

}

func importJokes(ctx context.Context, client *ent.Client) error {
	fmt.Println("import jokes")
	jokesBytes, err := os.ReadFile("seed/jokes.json")
	io.FatalIfError(err)

	type entry struct {
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
			SetBody(entry.Body).
			SetCategory(entry.Category).
			Save(ctx)

		if err != nil {
			fmt.Println(entry.Body[:10], err)
			continue
		}
	}
	return nil
}

func importBbs(ctx context.Context, client *ent.Client) error {
	fmt.Println("import bbs hosts")
	bbsBytes, err := os.ReadFile("seed/bbs.json")
	io.FatalIfError(err)

	type entry struct {
		Phone      string `json:"phone"`
		Names      string `json:"names"`
		Location   string `json:"location"`
		Sysop      string `json:"sysop"`
		SystemName string `json:"system_name"`
		YearFrom   int    `json:"year_from"`
		YearTo     int    `json:"year_to"`
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

		systemName := schema.HostName(entry.SystemName)
		_, err := client.Host.Create().
			SetType(host.TypeBbs).
			SetName(systemName).
			SetOrganization(names[0]).
			SetContact(sysops[0]).
			SetPhone([]schema.PhoneNumber{schema.PhoneNumber(entry.Phone)}).
			SetLocation(entry.Location).
			Save(ctx)

		if err != nil {
			fmt.Println(systemName, err)
			continue
		}

	}
	return nil
}

func importServices(ctx context.Context, client *ent.Client) error {
	fmt.Println("import services")
	uumapBytes, err := os.ReadFile("seed/services.json")
	if err != nil {
		return err
	}

	var repr []ent.TcpService
	err = json.Unmarshal(uumapBytes, &repr)
	if err != nil {
		return err
	}

	for _, v := range repr {
		_, err := client.TcpService.Create().
			SetDescription(v.Description).
			SetName(v.Name).
			SetPort(v.Port).
			Save(ctx)
		if err != nil {
			fmt.Println(v.Name, err)
			continue
		}
	}

	return nil
}

func importUumap(ctx context.Context, client *ent.Client) error {
	fmt.Println("import umap")
	uumapBytes, err := os.ReadFile("seed/uumap.json")
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

		// add random services
		addRandomServices(ctx, client, h)

	}
	return nil
}

func addRandomServices(ctx context.Context, client *ent.Client, host *ent.Host) error {
	services := make([]*ent.TcpService, 0)
	services = append(services, client.TcpService.Query().Where(tcpservice.PortEQ(21)).OnlyX(ctx))
	services = append(services, client.TcpService.Query().Where(tcpservice.PortEQ(23)).OnlyX(ctx))
	services = append(services, client.TcpService.Query().Where(tcpservice.PortEQ(79)).OnlyX(ctx))
	services = append(services, client.TcpService.Query().Where(tcpservice.PortEQ(513)).OnlyX(ctx))

	c := 2 + rand.Intn(2)

	for i := 0; i < c; i++ {
		service := client.TcpService.Query().
			Where(tcpservice.PortLT(1024)).
			Order(func(s *sql.Selector) { s.OrderBy("RANDOM()") }).
			Limit(1).
			AllX(ctx)[0]

		contains := slices.Any(
			services,
			func(s *ent.TcpService) bool { return s.Port == service.Port },
		)

		if !contains {
			services = append(services, service)
		}
	}

	host.Update().AddServices(services...).ExecX(ctx)
	return nil
}

func importCsokavar(ctx context.Context, client *ent.Client) error {
	fmt.Println("import csokavar")
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

func importMilHosts(ctx context.Context, client *ent.Client) error {
	fmt.Println("import mil hosts")
	input, err := os.ReadFile("seed/mil.json")
	if err != nil {
		return err
	}

	type entry struct {
		Organization string `json:"organization"`
		Location     string `json:"location"`
		SystemName   string `json:"system_name"`
	}

	var repr []entry
	err = json.Unmarshal(input, &repr)
	if err != nil {
		return err
	}

	for _, v := range repr {
		// create a fake number
		st := "+1 808"
		st += fmt.Sprintf("%d", rand.Intn(9)+1)

		for i := 0; i < 6; i++ {
			st += fmt.Sprintf("%d", rand.Intn(10))
		}
		phone, err := schema.ParsePhoneNumber(st)
		io.FatalIfError(err)

		u := client.VirtualUser.Create().
			SetUser(milnet.GenerateUserId()).
			SetPassword(milnet.GenerateAccessCode()).
			SaveX(ctx)

		_, err = client.Host.Create().
			SetName(schema.HostName(v.SystemName)).
			SetType(host.TypeMil).
			SetOrganization(v.Organization).
			SetPhone([]schema.PhoneNumber{phone}).
			SetLocation(v.Location).
			AddVirtualusers(u).
			Save(ctx)

		if err != nil {
			fmt.Println(v.SystemName, err)
			continue
		}

	}
	return nil
}
