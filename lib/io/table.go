package io

import (
	"strings"

	"github.com/encse/altnet/lib/slices"
)

func Table(lines ...[]string) string {
	lines = flatten(lines)

	var colWidth []int
	for _, columns := range lines {
		for icol, col := range columns {
			if icol >= len(colWidth) {
				colWidth = append(colWidth, 0)
			}
			if len(col) > colWidth[icol] {
				colWidth[icol] = len(col)
			}
		}
	}
	var colsep = 2
	var res []string
	for _, columns := range lines {
		line := ""
		for icol, col := range columns {
			line += col + strings.Repeat(" ", colWidth[icol]-len(col)+colsep)
		}
		res = append(res, line)
	}
	return strings.Join(res, "\n") + "\n"
}

func flatten(input [][]string) [][]string {

	output := make([][]string, 0)

	for _, row := range input {
		lines := slices.Map(
			row, func(col string) []string { return strings.Split(col, "\n") })
		
		height := slices.Max(
			slices.Map(lines, func(col []string) int { return len(col) }))

		for j := 0; j < height; j++ {
			output = append(
				output,
				slices.Map(lines, func(col []string) string {
					if j < len(col) {
						return col[j]
					} else {
						return ""
					}
				}),
			)
		}
	}

	return output
}
