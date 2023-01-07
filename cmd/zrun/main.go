package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/encse/altnet/lib/config"
	"github.com/encse/altnet/lib/csokavar"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/log"
	"golang.org/x/exp/maps"
)

const gamesDir = "data/doors"

func main() {
	conf := config.Get()
	if conf.Dfrotz.Location == "" {
		fmt.Println("Zrun config is missing.")
		return
	}

	files, err := ioutil.ReadDir(gamesDir)
	if err != nil {
		log.Error(err)
		fmt.Println("Games cannot be found.")
		return
	}

	games := map[string]string{}
	for _, file := range files {
		extension := filepath.Ext(file.Name())
		if !file.IsDir() && extension == ".z5" {
			nickname := strings.TrimSuffix(file.Name(), extension)
			games[nickname] = path.Join(gamesDir, file.Name())
		}
	}

	name, err := io.ReadArgFromList("game", os.Args, 1, maps.Keys(games))
	io.FatalIfError(err)
	csokavar.RunCommand(conf.Dfrotz.Location, "-q", "-R", "/tmp", games[name])
}
