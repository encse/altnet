package phonenumbers

import (
	"strings"

	lib "github.com/nyaruka/phonenumbers"
)

type PhoneNumber string

func (pn PhoneNumber) Prev() (PhoneNumber, bool) {
	repr, err := lib.Parse(string(pn), "US")
	if err != nil {
		return pn, false
	}
	if repr.NationalNumber == nil {
		return pn, false
	}
	*repr.NationalNumber -= 1

	return PhoneNumber(lib.Format(repr, lib.INTERNATIONAL)), true
}

func (pn PhoneNumber) Next() (PhoneNumber, bool) {
	repr, err := lib.Parse(string(pn), "US")
	if err != nil {
		return pn, false
	}
	if repr.NationalNumber == nil {
		return pn, false
	}
	*repr.NationalNumber += 1

	return PhoneNumber(lib.Format(repr, lib.INTERNATIONAL)), true
}

func (pn PhoneNumber) ToAtdtString() (string, error) {
	repr, err := lib.Parse(string(pn), "US")
	if err != nil {
		return "", err
	}
	atdt := lib.Format(repr, lib.INTERNATIONAL)
	atdt = strings.ReplaceAll(atdt, " ext. ", ",") // add 2 seconds wait when calling extensions
	if strings.HasPrefix(atdt, "+1 ") {
		atdt = atdt[2:] // local number
	}
	atdt = strings.ReplaceAll(atdt, "+", "011") // international call
	atdt = strings.ReplaceAll(atdt, "-", "")
	atdt = strings.ReplaceAll(atdt, " ", "")
	return "ATDT" + atdt, nil
}

func ParsePhoneNumber(st string) (PhoneNumber, error) {
	repr, err := lib.Parse(st, "US")
	if err != nil {
		return PhoneNumber(""), err
	}
	return PhoneNumber(lib.Format(repr, lib.INTERNATIONAL)), nil
}

func ParsePhoneNumberSkipExtension(st string) (PhoneNumber, error) {

	repr, err := lib.Parse(st, "US")
	if err != nil {
		return PhoneNumber(""), err
	}
	repr.Extension = nil
	return PhoneNumber(lib.Format(repr, lib.INTERNATIONAL)), nil
}
