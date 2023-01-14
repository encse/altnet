package io

import "strings"

func Table(lines ...[]string) string {
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

	var res []string
	for _, columns := range lines {
		line := ""
		for icol, col := range columns {
			line += col + strings.Repeat(" ", colWidth[icol]-len(col)+1)
		}
		res = append(res, line)
	}
	return strings.Join(res, "\n")
}
