package inputs

import (
	"strconv"
	"strings"
)

func ReadIntSlice(line string) ([]int, error) {
	line = trimSquareBrackets(line)
	numbers := strings.Split(line, ",")
	intSlice := make([]int, 0, len(numbers))
	for _, number := range numbers {
		number = strings.TrimSpace(number)
		i, err := strconv.Atoi(number)
		if err != nil {
			return nil, err
		}
		intSlice = append(intSlice, i)
	}
	return intSlice, nil
}

func trimSquareBrackets(line string) string {
	line = strings.TrimSpace(line)
	for len(line) > 0 && line[0] == '[' && line[len(line)-1] == ']' {
		line = line[1 : len(line)-1]
		line = strings.TrimSpace(line)
	}
	return line
}
