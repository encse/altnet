package text

import (
	"strings"
)

func Linebreak(text string, width int) string {
	lines := strings.Split(text, "\n")
	for i := 0; i < len(lines); i++ {
		line := []rune(lines[i])
		ichSpace := 0
		nonEscapedChars := 0
		for ich := 0; ich < len(line); ich++ {
			nonEscapedChars++
			if line[ich] == ' ' {
				ichSpace = ich
			}
			if nonEscapedChars > width {
				if ichSpace > 0 {
					lines = append(lines, "")
					copy(lines[i+1:], lines[i:])
					lines[i] = strings.TrimRight(string(line[:ichSpace]), " ")
					lines[i+1] = strings.TrimRight(string(line[ichSpace+1:]), " ")
				}
				break
			}
		}
	}
	return strings.Join(lines, "\n")
}
