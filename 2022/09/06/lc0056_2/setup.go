package main

import (
	"fmt"
	"strconv"
	"strings"
)

func readIntervals(input string) [][]int {
	input = strings.TrimLeft(input, "[")
	input = strings.TrimRight(input, "]")
	parts := strings.Split(input, "],[")
	intervals := make([][]int, 0, len(parts))
	for _, part := range parts {
		pair := strings.Split(part, ",")
		l, err := strconv.Atoi(pair[0])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(pair[1])
		if err != nil {
			panic(err)
		}
		intervals = append(intervals, []int{l, r})
	}
	return intervals
}

func formatIntervals(intervals [][]int) string {
	parts := make([]string, 0, len(intervals))
	for _, interval := range intervals {
		parts = append(parts, fmt.Sprintf("[%d,%d]", interval[0], interval[1]))
	}
	return fmt.Sprintf("[%s]", strings.Join(parts, ","))
}
