package main

import (
	"fmt"
	"testing"
)

func TestNumIslands(t *testing.T) {
	tests := []struct {
		grid string
		want int
	}{
		{
			grid: `[
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]`,
			want: 1,
		},
		{
			grid: `[
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
			grid := readGrid(tt.grid)
			got := numIslands(grid)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
