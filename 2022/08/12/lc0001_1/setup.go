package lc0001_1

import (
	"fmt"
	"strconv"
	"strings"
)

func readInput(input string) []int {
	input = strings.TrimSpace(input)
	input = strings.TrimLeft(input, "[")
	input = strings.TrimRight(input, "]")
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

func formatInts(result []int) string {
	res := make([]string, 0, len(result))
	for _, i := range result {
		res = append(res, strconv.Itoa(i))
	}
	return fmt.Sprintf("[%s]", strings.Join(res, ","))
}
