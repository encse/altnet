package altnet

import (
	"context"
	"errors"
	"fmt"
	"hash/fnv"
	"io/fs"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/encse/altnet/ent/host"
	"github.com/encse/altnet/ent/schema"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/log"
	"github.com/encse/altnet/lib/uumap"
)

const altnetRoot = "data/altnet"
const seedRoot = "data/textfiles"

type FileInfo struct {
	fsFileInfo fs.FileInfo
	realPath   string
}

func (fi FileInfo) Name() string       { return fi.fsFileInfo.Name() }
func (fi FileInfo) Size() int64        { return fi.fsFileInfo.Size() }
func (fi FileInfo) Mode() fs.FileMode  { return fi.fsFileInfo.Mode() }
func (fi FileInfo) ModTime() time.Time { return fi.fsFileInfo.ModTime() }
func (fi FileInfo) IsDir() bool        { return fi.fsFileInfo.IsDir() }
func (fi FileInfo) Sys() any           { return fi }
func (fi FileInfo) RealPath() string   { return fi.realPath }

func Files(ctx context.Context) ([]FileInfo, error) {
	host, err := GetHost(ctx)
	if err != nil {
		return nil, err
	}
	user, err := GetUser(ctx)
	if err != nil {
		return nil, err
	}

	seedHost(ctx, host)

	res := make([]FileInfo, 0)

	dirs := []string{
		getAltnetUserDir(host, user),
		getAltnetSystemDir(host),
		getAltnetSeedDir(host),
	}

	for _, dir := range dirs {
		files, err := ioutil.ReadDir(dir)
		if errors.Is(err, fs.ErrNotExist) {
			continue
		} else if err != nil {
			log.Error(err)
			continue
		}

		for _, file := range files {
			// resolve symlinks
			fi, err := os.Stat(path.Join(dir, file.Name()))
			if err != nil {
				log.Error()
				continue
			}
			res = append(res,
				FileInfo{
					fsFileInfo: fi,
					realPath:   path.Join(dir, file.Name()),
				})
		}
	}

	return res, nil
}

func GetFileInfo(ctx context.Context, name string) (FileInfo, error) {
	files, err := Files(ctx)
	if err != nil {
		return FileInfo{}, err
	}

	for _, fi := range files {
		if fi.Name() == name {
			return fi, nil
		}
	}

	return FileInfo{}, io.UserFriendlyError{Err: fs.ErrNotExist}
}

func Open(ctx context.Context, name string) (*os.File, error) {
	fi, err := GetFileInfo(ctx, name)
	if err != nil {
		return nil, err
	}
	return os.Open(fi.RealPath())
}

func Cat(ctx context.Context, fi FileInfo) error {
	RunHiddenCommandWithStdErrRedirectedToStdout(ctx, "/bin/cat", fi.RealPath())
	return nil
}

func More(ctx context.Context, fi FileInfo) error {
	RunHiddenCommandWithStdErrRedirectedToStdout(ctx, "/bin/more", fi.RealPath())
	return nil
}

func getAltnetUserDir(host schema.HostName, user User) string {
	return path.Join(altnetRoot, string(host), "usr", string(user))
}
func getAltnetSeedDir(host schema.HostName) string {
	return path.Join(altnetRoot, string(host), "seed")
}
func getAltnetSystemDir(host schema.HostName) string {
	return path.Join(altnetRoot, string(host), "sys")
}

func seedHost(ctx context.Context, hostName schema.HostName) error {
	targetDir := getAltnetSeedDir(hostName)

	targetDir, err := filepath.Abs(targetDir)
	if err != nil {
		return err
	}

	exists, err := fileExists(targetDir)
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	err = os.MkdirAll(targetDir, 0777)
	if err != nil {
		return err
	}

	if hostName != "csokavar" {
		err = addRandomFiles(hostName, targetDir)
		if err != nil {
			log.Warn(fmt.Errorf("could not add random files, %v", err))
		}
	}

	err = createBbsList(ctx, targetDir)
	if err != nil {
		log.Warn(fmt.Errorf("could not create bbslist.txt, %v", err))
	}

	return nil
}

func addRandomFiles(hostName schema.HostName, targetDir string) error {
	seedFiles, err := getAllFiles(seedRoot)
	if err != nil {
		return err
	}

	r := newRand(string(hostName))
	batchCount := r.Int()%5 + 2

	for i := 0; i < batchCount; i++ {
		fileCount := r.Int()%10 + 5
		ptr := r.Int()
		for j := 0; j < fileCount; j++ {
			seedPath := seedFiles[ptr%len(seedFiles)]
			ptr++
			srcPath, err := filepath.Abs(seedPath)
			if err != nil {
				return err
			}
			srcPath, err = filepath.Rel(targetDir, srcPath)
			if err != nil {
				return err
			}
			err = os.Symlink(srcPath, path.Join(targetDir, path.Base(seedPath)))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// createBbsList randomly picks 20 bbs-es from the hosts database
func createBbsList(ctx context.Context, targetDir string) error {
	network, err := uumap.NetworkConn()

	if err != nil {
		return err
	}
	defer network.Close()
	hosts, err := network.Client.Host.
		Query().
		Where(
			host.TypeEQ(host.TypeBbs),
		).
		Order(func(s *sql.Selector) {
			s.OrderBy("RANDOM()")
		}).
		Limit(20).
		All(ctx)

	if err != nil {
		return err
	}

	lines := make([][]string, 0, len(hosts))
	for i, host := range hosts {
		if len(host.Phone) > 0 {
			lines = append(lines, []string{
				fmt.Sprintf("%02d.", i+1),
				host.Organization,
				host.Phone[0].ToUsLocalString(),
			})
		}
	}

	st := io.Table(lines...)
	return os.WriteFile(
		path.Join(targetDir, "bbslist.txt"),
		[]byte(st),
		0644,
	)
}

func fileExists(file string) (bool, error) {
	if _, err := os.Stat(file); err == nil {
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return false, err
	}
}

func getAllFiles(dir string) ([]string, error) {
	seedPaths := make([]string, 0, 10000)

	// collect files under seedRoot
	err := filepath.WalkDir(seedRoot, func(file string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			seedPaths = append(seedPaths, file)
		}

		return nil
	})
	return seedPaths, err
}

func newRand(seed string) *rand.Rand {
	h := fnv.New32a()
	h.Write([]byte(seed))
	h.Sum32()
	return rand.New(rand.NewSource(int64(h.Sum32())))
}
