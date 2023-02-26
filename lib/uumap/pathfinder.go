package uumap

import (
	"container/list"
	"context"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/encse/altnet/ent/schema"

	_ "github.com/mattn/go-sqlite3"
)

func FindPaths(
	ctx context.Context,
	network Network,
	sourceHost schema.HostName,
	targetHost schema.HostName,
	maxCount int,
) [][]schema.HostName {
	res := make([][]schema.HostName, 0)
	q := list.New()
	q.PushBack([]schema.HostName{sourceHost})
	seen := mapset.NewSet[schema.HostName]()

	for len(res) < maxCount && q.Len() > 0 {
		path := q.Front().Value.([]schema.HostName)
		q.Remove(q.Front())

		hostName := path[len(path)-1]
		seen.Add(schema.HostName(hostName))

		if hostName == targetHost {
			res = append(res, path)
		}

		host, _ := network.Lookup(ctx, hostName)
		if host != nil {
			for _, hostNext := range host.Neighbours {
				if !seen.Contains(hostNext) {
					res := make([]schema.HostName, len(path), len(path)+1)
					copy(res, path)
					res = append(res, hostNext)
					q.PushBack(res)
				}
			}
		}
	}

	return res
}
