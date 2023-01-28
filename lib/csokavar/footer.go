package csokavar

import (
	"fmt"
	"io/ioutil"

	"github.com/encse/altnet/lib/io"
)

func Footer(screenWidth int) (string, error) {
	footer, err := ioutil.ReadFile("data/altnet/csokavar/sys/footer.txt")
	if err != nil {
		return "", fmt.Errorf("Could not get footer, %w", err)
	}
	return io.Center(string(footer)+"\n", screenWidth), nil
}
