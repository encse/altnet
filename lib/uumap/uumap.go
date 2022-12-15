package uumap

import (
	"encoding/json"
	"io/ioutil"
	"sort"

	"github.com/encse/altnet/lib/tools"
)

type Uunode = struct {
	Entry string   `json:"entry"`
	Hosts []string `json:"hosts"`
}

type Uumap = map[string]Uunode

func GetUumap() (Uumap, error) {
	uumapBytes, err := ioutil.ReadFile("data/uumap.json")
	if err != nil {
		return Uumap{}, err
	}

	var uumap Uumap
	err = json.Unmarshal(uumapBytes, &uumap)
	return uumap, err
}

func FindPaths(network Uumap, sourceHost string, targetHost string, maxLength int, maxCount int) [][]string {
	res := make([][]string, 0)
	findPaths(network, targetHost, maxLength, []string{}, sourceHost, &res)
	sort.SliceStable(res, func(i, j int) bool {
		return len(res[i]) < len(res[j])
	})

	if len(res) > maxCount {
		return res[:maxCount]
	} else {
		return res
	}
}

func findPaths(network Uumap, targetHost string, maxLength int, path []string, host string, allPaths *([][]string)) {
	if tools.Contains(path, host) || len(path) > maxLength {
		return
	}
	path = append(path, host)
	if host == targetHost {
		*allPaths = append(*allPaths, path)
	} else if entry, ok := network[host]; ok {
		for _, hostNext := range entry.Hosts {
			pathNext := make([]string, len(path), len(path)+1)
			copy(pathNext, path)
			findPaths(network, targetHost, maxLength, pathNext, hostNext, allPaths)
		}
	}
}
