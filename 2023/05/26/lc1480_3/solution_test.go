package lc1480_3

import (
	"fmt"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
	"github.com/dowenliu-xyz/leetcode/internal/outputs"
)

func Test_runningSum(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `[1,2,3,4]`,
			want:  `[1,3,6,10]`,
		},
		{
			input: `[1,1,1,1,1]`,
			want:  `[1,2,3,4,5]`,
		},
		{
			input: `[3,1,2,10,1]`,
			want:  `[3,4,6,16,17]`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			nums, err := inputs.ReadIntSlice(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			got := runningSum(nums)
			want := strings.ReplaceAll(tt.want, " ", "")
			if got := outputs.SprintIntSlice(got); got != want {
				t.Errorf("runningSum(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
