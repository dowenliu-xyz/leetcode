package lc0136_1

import (
	"fmt"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func Test_singleNumber(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `[2,2,1]`,
			want:  1,
		},
		{
			input: `[4,1,2,1,2]`,
			want:  4,
		},
		{
			input: `[1]`,
			want:  1,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			nums, err := inputs.ReadIntSlice(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			if got := singleNumber(nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
