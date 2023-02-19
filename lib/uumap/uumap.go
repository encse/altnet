package uumap

import (
	"container/list"
	"encoding/json"
	"io/ioutil"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/phonenumbers"
)

type Uunode = struct {
	Entry          string      `json:"entry"`
	Hosts          []string    `json:"neighbours"`
	Country        string      `json:"country"`
	HostName       altnet.Host `json:"system-name"`
	MachineType    string      `json:"machine-type"`
	Organization   string      `json:"organization"`
	Contact        string      `json:"contact"`
	ContactAddress string      `json:"contact-address"`
	Phone          []string    `json:"phone"`
	Location       string      `json:"location"`
	GeoLocation    string      `json:"geo-location"`
}

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

type Phonebook struct {
	items map[phonenumbers.PhoneNumber]altnet.Host
}

// Lookup checks the number in the phonebook and returns with a hostname if found.
// If the host doesn't need an extension but an extension is provided in the phone number
// we return with the host regardless. However if the host requires an extension and no
// extension is provided in the phone number we return with failure.
// This is analogous to dialing a number: if the extension is not needed, it is simply
// ignored.
func (phonebook Phonebook) Lookup(phoneNumber phonenumbers.PhoneNumber) (altnet.Host, bool) {
	if host, ok := phonebook.items[phoneNumber]; ok {
		return host, true
	}

	withoutExt, err := phonenumbers.ParsePhoneNumberSkipExtension(string(phoneNumber))
	if err != nil {
		return altnet.Host(""), false
	}
	host, ok := phonebook.items[withoutExt]
	return host, ok
}

func GetPhonebook() (Phonebook, error) {
	phonebookBytes, err := ioutil.ReadFile("data/phonebook.json")
	if err != nil {
		return Phonebook{}, err
	}

	var phonebook Phonebook
	err = json.Unmarshal(phonebookBytes, &phonebook.items)
	return phonebook, err
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
