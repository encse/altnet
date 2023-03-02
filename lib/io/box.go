package io

import (
	"fmt"
	"strings"

	"github.com/encse/altnet/lib/slices"
)

func Box(lines ...string) string {
	maxLength := slices.Max(slices.Map(lines, func(line string) int { return len(line) }))
	horizontalLine := strings.Repeat("*", maxLength+4)

	res := horizontalLine + "\n"
	for _, line := range lines {
		res += fmt.Sprintf("* %s *\n", Center(line, maxLength))
	}
	res += horizontalLine + "\n"
	return res
}
