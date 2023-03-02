package io

import (
	"strings"
	"unicode/utf8"
)

func Center(st string, width int) string {
	lines := strings.Split(st, "\n")
	maxWidth := 0
	for _, line := range lines {
		if utf8.RuneCountInString(line) >= maxWidth {
			maxWidth = utf8.RuneCountInString(line)
		}
	}
	for i, line := range lines {
		padLeft := (width - maxWidth) / 2
		padRight := width - maxWidth - padLeft
		if padLeft <= 0 {
			continue
		}
		lines[i] = strings.Repeat(" ", padLeft) + line + strings.Repeat(" ", padRight)
	}
	return strings.Join(lines, "\n")
}
