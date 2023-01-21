package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/log"
	"github.com/encse/altnet/lib/uumap"
)

func main() {
	ctx := altnet.ContextFromEnv(context.Background())
	host, err := altnet.GetHost(ctx)
	io.FatalIfError(err)

	targetHost, err := io.ReadArg("host", os.Args, 1)
	io.FatalIfError(err)

	targetHost = strings.ToLower(targetHost)

	log.Info("loading nodes")
	network, err := uumap.GetUumap()
	io.FatalIfError(err)

	if _, ok := network[targetHost]; !ok {
		fmt.Println("unknown host", targetHost)
		return
	}

	log.Info("finding paths")
	res := uumap.FindPaths(network, string(host), targetHost, 3)

	if len(res) == 0 {
		fmt.Println("no path to host")
		return
	}
	fmt.Println("collecting edges")
	sb := strings.Builder{}

	for edge := range getEdges(res, string(host), targetHost).Iter() {
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

func getEdges(paths [][]string, startHost string, targetHost string) mapset.Set[string] {
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

			edge := fmt.Sprintf("[ %s ] %s -> [ %s ] %s\n", path[j-1], style1, path[j], style2)
			if !res.Contains(edge) {
				res.Add(edge)
			}
		}
	}
	return res
}
