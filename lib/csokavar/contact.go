package csokavar

import (
	"fmt"
	"io/ioutil"

	"github.com/encse/altnet/lib/io"
)

func GpgKey(screenWidth int) (string, error) {
	key, err := ioutil.ReadFile("data/altnet/csokavar/sys/encse.gpg")
	if err != nil {
		return "", fmt.Errorf("Could not get gpg key, %w", err)
	}
	res := "Gpg key, reach me at encse@csokavar.hu\n"
	res += io.Center(string(key)+"\n", screenWidth)
	return res, nil
}
