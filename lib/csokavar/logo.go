package csokavar

import (
	"fmt"
	"io/ioutil"

	"github.com/encse/altnet/lib/io"
)

func Logo(screenWidth int) (string, error) {
	logo, err := ioutil.ReadFile("data/altnet/csokavar/sys/logo.txt")
	if err != nil {
		return "", fmt.Errorf("Could not get logo, %w", err)
	}
	return io.Center(string(logo), screenWidth) + "\n", nil
}
