package main

import (
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
