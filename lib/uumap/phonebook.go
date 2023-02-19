package uumap

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/encse/altnet/lib/maps"
	"github.com/encse/altnet/lib/phonenumbers"
	"github.com/encse/altnet/lib/slices"
)

type Phonebook struct {
	items map[phonenumbers.PhoneNumber]Host
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

// Lookup checks the number in the phonebook and returns with a hostname if found.
// If the host doesn't need an extension but an extension is provided in the phone number
// we return with the host regardless. However if the host requires an extension and no
// extension is provided in the phone number we return with failure.
// This is analogous to dialing a number: if the extension is not needed, it is simply
// ignored.
func (phonebook Phonebook) Lookup(phoneNumber phonenumbers.PhoneNumber) (Host, bool) {
	if host, ok := phonebook.items[phoneNumber]; ok {
		return host, true
	}

	withoutExt, err := phonenumbers.ParsePhoneNumberSkipExtension(string(phoneNumber))
	if err != nil {
		return Host(""), false
	}
	host, ok := phonebook.items[withoutExt]
	return host, ok
}

func (phonebook Phonebook) LookupByPrefix(prefix string) []phonenumbers.PhoneNumber {
	return slices.Filter(maps.Keys(phonebook.items),
		func(phoneNumber phonenumbers.PhoneNumber) bool {
			return strings.HasPrefix(string(phoneNumber), prefix)
		})
}
