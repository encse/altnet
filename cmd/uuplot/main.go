package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/uumap"
	log "github.com/sirupsen/logrus"
)

func main() {

	targetHost, err := io.ReadArg("host", os.Args, 1)
	io.FatalIfError(err)

	targetHost = strings.ToLower(targetHost)

	network, err := uumap.GetUumap()
	io.FatalIfError(err)

	network["home"] = uumap.Uunode{
		Entry: "",
		Hosts: []string{
			"adaptex",
			"aemsrc",
			"attvcr",
			"bpsm",
			"flux",
			"grex",
			"mimsy",
			"oddjob",
			"oracle",
			"tandem",
			"veritas",
		},
	}

	startHost := "home"

	res := uumap.FindPaths(network, startHost, targetHost, 5, 6)

	if len(res) == 0 {
		fmt.Println("no path to host")
		return
	}

	sb := strings.Builder{}

	for edge := range getEdges(res, startHost, targetHost).Iter() {
		sb.WriteString(edge)
	}

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
