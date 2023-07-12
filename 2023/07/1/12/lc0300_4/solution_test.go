package lc0300

import (
	"fmt"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func Test_lengthOfLIS(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `[10,9,2,5,3,7,101,18]`,
			want:  4,
		},
		{
			input: `[0,1,0,3,2,3]`,
			want:  4,
		},
		{
			input: `[7,7,7,7,7,7,7]`,
			want:  1,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			nums, err := inputs.ReadIntSlice(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			got := lengthOfLIS(nums)
			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
