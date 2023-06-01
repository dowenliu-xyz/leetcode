package inputs

import (
	"fmt"
	"strconv"
	"strings"
)

func ReadIntSlice(line string) ([]int, error) {
	line = strings.ReplaceAll(line, " ", "")
	line = strings.ReplaceAll(line, "\n", "")
	line = trimSquareBrackets(line)
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
	}
	return line
}

func ReadIntMatrix(input string) ([][]int, error) {
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ReplaceAll(input, "\n", "")
	if input == "" || input == "[]" {
		return nil, nil
	}
	if !strings.HasPrefix(input, "[[") || !strings.HasSuffix(input, "]]") {
		return nil, fmt.Errorf("invalid input: %s", input)
	}
	input = strings.Trim(input, "[]")
	rows := strings.Split(input, "],[")
	matrix := make([][]int, len(rows))
	for i, row := range rows {
		elements := strings.Split(row, ",")
		matrix[i] = make([]int, len(elements))
		for j, element := range elements {
			num, err := strconv.Atoi(element)
			if err != nil {
				return nil, fmt.Errorf("invalid input: %s", input)
			}
			matrix[i][j] = num
		}
	}
	return matrix, nil
}
