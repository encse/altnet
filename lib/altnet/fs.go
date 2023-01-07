package altnet

import (
	"context"
	"io/fs"
	"io/ioutil"
	"path"
)

const altnetRoot = "data/altnet"

func Files(ctx context.Context) ([]fs.FileInfo, error) {
	host, err := GetHost(ctx)
	if err != nil {
		return nil, err
	}
	files, err := ioutil.ReadDir(path.Join(altnetRoot, string(host)))
	if err != nil {
		return nil, err
	}
	return files, nil
}
