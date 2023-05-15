package inputs

import (
	"fmt"
	"strconv"
	"strings"
)

func ReadIntSlice(line string) ([]int, error) {
	line = trimSquareBrackets(line)
	line = strings.TrimSpace(line)
	if line == "" {
		return nil, nil
	}
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

func ReadIntSquareInLine(input string) ([][]int, error) {
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, " ", "")
	// [[1,2,3],[4,5,6],[7,8,9]]
	if input == "" {
		return nil, nil
	}
	if !strings.HasPrefix(input, "[[") || !strings.HasSuffix(input, "]]") {
		return nil, fmt.Errorf("invalid input: %s", input)
	}
	input = input[2 : len(input)-2]
	// 1,2,3],[4,5,6],[7,8,9
	lines := strings.Split(input, "],[")
	// 1,2,3 | 4,5,6 | 7,8,9
	matrix := make([][]int, 0, len(lines))
	for _, line := range lines {
		intSlice, err := ReadIntSlice(line)
		if err != nil {
			return nil, err
		}
		matrix = append(matrix, intSlice)
	}
	return matrix, nil
}
