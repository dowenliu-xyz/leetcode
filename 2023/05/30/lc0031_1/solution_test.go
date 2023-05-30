package lc0031_1

import (
	"fmt"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
	"github.com/dowenliu-xyz/leetcode/internal/outputs"
)

func Test_nextPermutation(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `[1,2,3]`,
			want:  `[1,3,2]`,
		},
		{
			input: `[3,2,1]`,
			want:  `[1,2,3]`,
		},
		{
			input: `[1,1,5]`,
			want:  `[1,5,1]`,
		},
		{
			input: `[1]`,
			want:  `[1]`,
		},
		{
			input: `[1,2]`,
			want:  `[2,1]`,
		},
		{
			input: `[2,3,1]`,
			want:  `[3,1,2]`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			nums, err := inputs.ReadIntSlice(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			nextPermutation(nums)
			got := outputs.SprintIntSlice(nums)
			if got != tt.want {
				t.Errorf("nextPermutation(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
