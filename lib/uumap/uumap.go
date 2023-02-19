package uumap

import (
	"container/list"
	"encoding/json"
	"io/ioutil"

	mapset "github.com/deckarep/golang-set/v2"
)

type Uunode = struct {
	Entry          string   `json:"entry"`
	Hosts          []string `json:"neighbours"`
	Country        string   `json:"country"`
	HostName       Host     `json:"system-name"`
	MachineType    string   `json:"machine-type"`
	Organization   string   `json:"organization"`
	Contact        string   `json:"contact"`
	ContactAddress string   `json:"contact-address"`
	Phone          []string `json:"phone"`
	Location       string   `json:"location"`
	GeoLocation    string   `json:"geo-location"`
}

type Host string
type Uumap = map[string]Uunode

func GetUumap() (Uumap, error) {
	uumapBytes, err := ioutil.ReadFile("data/uumap.json")
	if err != nil {
		return Uumap{}, err
	}

	var uumap Uumap
	err = json.Unmarshal(uumapBytes, &uumap)

	uumap["csokavar"] = Uunode{
		Hosts: []string{
			"oddjob",
			"adaptex",
			"aemsrc",
			"bpsm",
			"tandem",
			"oracle",
			"veritas",
			"mimsy",
		},
	}
	return uumap, err
}

func FindPaths(network Uumap, sourceHost string, targetHost string, maxCount int) [][]string {
	res := make([][]string, 0)
	q := list.New()
	q.PushBack([]string{sourceHost})
	seen := mapset.NewSet[string]()

	for len(res) < maxCount && q.Len() > 0 {
		path := q.Front().Value.([]string)
		q.Remove(q.Front())

		host := path[len(path)-1]
		seen.Add(host)

		if host == targetHost {
			res = append(res, path)
		} else if entry, ok := network[host]; ok {
			for _, hostNext := range entry.Hosts {
				if !seen.Contains(hostNext) {
					res := make([]string, len(path), len(path)+1)
					copy(res, path)
					res = append(res, hostNext)
					q.PushBack(res)
				}
			}
		}
	}

	return res
}
