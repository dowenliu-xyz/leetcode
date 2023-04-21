package inputs

import "strings"

func ReadNoneBlankLines(input string) []string {
	inputLines := strings.Split(input, "\n")
	lines := make([]string, 0, len(inputLines))
	for _, line := range inputLines {
		line = strings.TrimSpace(line)
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}
