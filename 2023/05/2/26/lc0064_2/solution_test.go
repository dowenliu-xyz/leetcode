package lc0064_2

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestMinPathSum(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "[[1,3,1],[1,5,1],[4,2,1]]",
			want:  7,
		},
		{
			input: "[[1,2,3],[4,5,6]]",
			want:  12,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			grid, err := readGrid(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			got := minPathSum(grid)
			if got != tt.want {
				t.Fatalf("got: %d, want: %d", got, tt.want)
			}
		})
	}
}

func readGrid(input string) ([][]int, error) {
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, " ", "")
	if !strings.HasPrefix(input, "[[") || !strings.HasSuffix(input, "]]") {
		return nil, fmt.Errorf("input should start with [[ and end with ]]")
	}
	input = strings.TrimPrefix(input, "[[")
	input = strings.TrimSuffix(input, "]]")
	rows := strings.Split(input, "],[")
	grid := make([][]int, len(rows))
	for i, row := range rows {
		cols := strings.Split(row, ",")
		grid[i] = make([]int, len(cols))
		for j, col := range cols {
			v, err := strconv.Atoi(col)
			if err != nil {
				return nil, err
			}
			grid[i][j] = v
		}
	}
	return grid, nil
}
