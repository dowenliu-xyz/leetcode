package lc0054_2

import (
	"fmt"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
	"github.com/dowenliu-xyz/leetcode/internal/outputs"
)

func Test_spiralOrder(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `
[[1,2,3],[4,5,6],[7,8,9]]
`,
			want: `[1,2,3,6,9,8,7,4,5]`,
		},
		{
			input: `
[[1,2,3,4],[5,6,7,8],[9,10,11,12]]
`,
			want: `[1,2,3,4,8,12,11,10,9,5,6,7]`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			matrix, err := inputs.ReadIntMatrix(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			got := spiralOrder(matrix)
			if got := outputs.SprintIntSlice(got); got != tt.want {
				t.Errorf("spiralOrder(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
