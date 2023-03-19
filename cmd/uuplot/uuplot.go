package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/encse/altnet/ent/host"
	"github.com/encse/altnet/ent/user"
	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/log"
	"github.com/encse/altnet/lib/uumap"
	"github.com/encse/altnet/schema"
)

func main() {
	ctx := altnet.ContextFromEnv(context.Background())

	realUser, err := altnet.GetRealUser(ctx)
	io.FatalIfError(err)

	currentHost, err := altnet.GetHost(ctx)
	io.FatalIfError(err)

	st, err := io.ReadArg("host", os.Args, 1)
	io.FatalIfError(err)

	targetHostName := schema.HostName(strings.ToLower(st))

	network, err := uumap.NetworkConn()
	io.FatalIfError(err)
	defer network.Close()

	targetHost, err := network.Lookup(ctx, targetHostName)
	io.FatalIfError(err)
	if targetHost == nil {
		fmt.Println("unknown host", targetHostName)
		return
	}

	log.Info("finding paths")
	res := uumap.FindPaths(ctx, network, currentHost, targetHostName, 3)

	if len(res) == 0 {
		fmt.Println("no path to host")
		return
	}
	fmt.Println("collecting edges")
	sb := strings.Builder{}

	hacked := make([]schema.HostName, 0)

	err = network.Client.User.Query().
		Where(user.UserEQ(realUser)).
		QueryHosts().
		Select(host.FieldName).
		Scan(ctx, &hacked)
	io.FatalIfError(err)

	for edge := range getEdges(res, currentHost, targetHostName, mapset.NewSet(hacked...)).Iter() {
		sb.WriteString(edge)
	}

	fmt.Println("rendering")
	cmd := exec.Command("/usr/bin/graph-easy")
	cmd.Stdin = strings.NewReader(sb.String())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Error(err)
	}
}

func getEdges(
	paths [][]schema.HostName,
	startHost schema.HostName,
	targetHost schema.HostName,
	hacked mapset.Set[schema.HostName],
) mapset.Set[string] {
	res := mapset.NewSet[string]()
	for _, path := range paths {
		for j := 1; j < len(path); j++ {
			style1 := ""
			if path[j-1] == targetHost || path[j-1] == startHost {
				style1 = "{ border: 1px dotted black; }"
			}

			style2 := ""
			if path[j] == targetHost || path[j] == startHost {
				style2 = "{ border: 1px dotted black; }"
			}
			from := string(path[j-1])
			if hacked.Contains(path[j-1]) {
				from += "*"
			}
			to := string(path[j])

			if hacked.Contains(path[j]) {
				to += "*"
			}
			edge := fmt.Sprintf("[ %s ] %s -> [ %s ] %s\n", from, style1, to, style2)
			if !res.Contains(edge) {
				res.Add(edge)
			}
		}
	}
	return res
}
