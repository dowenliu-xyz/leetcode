package main

import (
	"fmt"
	"strconv"
	"strings"
)

func readMatrix(input string) [][]int {
	input = strings.TrimLeft(input, "[")
	input = strings.TrimRight(input, "]")
	rows := strings.Split(input, "],[")
	matrix := make([][]int, 0, len(rows))
	for _, row := range rows {
		nums := strings.Split(row, ",")
		mRow := make([]int, 0, len(row))
		for _, num := range nums {
			i, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			mRow = append(mRow, i)
		}
		matrix = append(matrix, mRow)
	}
	return matrix
}

func formatInts(ints []int) string {
	strs := make([]string, 0, len(ints))
	for _, i := range ints {
		strs = append(strs, strconv.Itoa(i))
	}
	return fmt.Sprintf("[%s]", strings.Join(strs, ","))
}
