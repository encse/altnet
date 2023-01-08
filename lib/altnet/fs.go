package altnet

import (
	"context"
	"io/fs"
	"io/ioutil"
	"os"
	"path"

	"github.com/encse/altnet/lib/io"
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

func Open(ctx context.Context, name string) (*os.File, error) {
	host, err := GetHost(ctx)
	if err != nil {
		return nil, err
	}

	dir := path.Join(altnetRoot, string(host))
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.Name() == name {
			return os.Open(path.Join(dir, file.Name()))
		}
	}
	return nil, &io.UserFriendlyError{Err: fs.ErrNotExist}
}
