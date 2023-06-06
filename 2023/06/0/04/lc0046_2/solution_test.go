package lc0046

import (
	"fmt"
	"sort"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
	"github.com/dowenliu-xyz/leetcode/internal/outputs"
)

func Test_permute(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `[1,2,3]`,
			want:  `[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]`,
		},
		{
			input: `[0,1]`,
			want:  `[[0,1],[1,0]]`,
		},
		{
			input: `[1]`,
			want:  `[[1]]`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			nums, err := inputs.ReadIntSlice(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			wantMatrix, err := inputs.ReadIntMatrix(tt.want)
			if err != nil {
				t.Fatal(err)
			}
			matrix := permute(nums)
			sortMatrix(wantMatrix)
			sortMatrix(matrix)
			got, want := outputs.SprintIntMatrix(matrix), outputs.SprintIntMatrix(wantMatrix)
			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}

func sortMatrix(matrix [][]int) {
	rows := len(matrix)
	if rows == 0 {
		return
	}
	cols := len(matrix[0])
	if cols == 0 {
		return
	}
	for _, row := range matrix {
		sort.Ints(row)
	}
	sort.SliceIsSorted(matrix, func(i, j int) bool {
		for col := 0; col < cols; col++ {
			if matrix[i][col] != matrix[j][col] {
				return matrix[i][col] < matrix[j][col]
			}
		}
		return false
	})
}
