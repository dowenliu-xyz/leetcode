package main

import (
	"fmt"
	"strconv"
	"strings"
)

func readInput(input string) []int {
	input = strings.TrimSpace(input)
	input = strings.TrimLeft(input, "[")
	input = strings.TrimRight(input, "]")
	if strings.TrimSpace(input) == "" {
		return nil
	}
	parts := strings.Split(input, ",")
	res := make([]int, 0, len(parts))
	for _, part := range parts {
		i, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
		res = append(res, i)
	}
	return res
}

func formatResult(result [][]int) string {
	lines := make([]string, 0, len(result))
	for _, vals := range result {
		line := make([]string, 0, len(vals))
		for _, val := range vals {
			line = append(line, strconv.Itoa(val))
		}
		lines = append(lines, fmt.Sprintf("[%s]", strings.Join(line, ",")))
	}
	return fmt.Sprintf("[%s]", strings.Join(lines, ","))
}
