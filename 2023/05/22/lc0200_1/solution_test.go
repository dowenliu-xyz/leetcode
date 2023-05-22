package lc0200_1

import (
	"fmt"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func Test_numIslands(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `[
["1","1","1","1","0"],
["1","1","0","1","0"],
["1","1","0","0","0"],
["0","0","0","0","0"]
]`,
			want: 1,
		},
		{
			input: `[
["1","1","0","0","0"],
["1","1","0","0","0"],
["0","0","1","0","0"],
["0","0","0","1","1"]
]`,
			want: 3,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			grid, err := readInput(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			got := numIslands(grid)
			if got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func readInput(input string) ([][]byte, error) {
	square, err := inputs.ReadStringSquare(input)
	if err != nil {
		return nil, err
	}
	grid := make([][]byte, len(square))
	for i := range square {
		row := square[i]
		grid[i] = make([]byte, len(row))
		for j, s := range row {
			switch s {
			case "1":
				grid[i][j] = '1'
			case "0":
				grid[i][j] = '0'
			default:
				return nil, fmt.Errorf("invalid input")
			}
		}
	}
	return grid, nil
}
