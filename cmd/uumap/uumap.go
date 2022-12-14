package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/encse/altnet/lib/io"
	log "github.com/sirupsen/logrus"
)

type Uumap = map[string]struct {
	Entry string   `json:"entry"`
	Hosts []string `json:"hosts"`
}

func main() {
	host := ""
	if len(os.Args) > 1 {
		host = os.Args[1]
	}

	for host == "" {
		var err error
		host, err = io.Readline("host: ")
		if err != nil {
			log.Fatal(err)
		}
		host = strings.ToLower(host)
	}

	uumap, err := GetUumap()
	if err != nil {
		log.Error(err)
		fmt.Println("couldn't load uumap database")
		return
	}
	if entry, ok := uumap[host]; ok {
		fmt.Println(entry.Entry)
	} else {
		fmt.Println("host not found")
	}

}

func GetUumap() (Uumap, error) {
	uumapBytes, err := ioutil.ReadFile("data/uumap.json")
	if err != nil {
		return Uumap{}, err
	}

	var uumap Uumap
	err = json.Unmarshal(uumapBytes, &uumap)
	return uumap, err
}

// func foo() {
// 	uumap, _ := GetUumap()

// 	hostsBytes, _ := ioutil.ReadFile("data/hosts")
// 	hosts := string(hostsBytes)
// 	res := Uumap{}
// 	for _, line := range strings.Split(hosts, "\n") {
// 		parts := strings.Split(strings.TrimSpace(line), " ")
// 		host := parts[0]
// 		if entry, ok := uumap[host]; ok {
// 			res[host] = entry
// 		} else {
// 			fmt.Println(host)
// 		}
// 	}

// 	bytes, err := json.MarshalIndent(res, "", "    ")
// 	if err != nil {
// 		log.Error("ERROR: fail to marshall json, %w", err.Error())
// 	}
// 	ioutil.WriteFile("data/uumap.json", bytes, 0666)
// }
