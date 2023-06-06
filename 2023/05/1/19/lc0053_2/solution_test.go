package lc0053_2

import (
	"fmt"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func Test_maxSubArray(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"[-2,1,-3,4,-1,2,1,-5,4]", 6},
		{"[1]", 1},
		{"[5,4,-1,7,8]", 23},
		{"[-1]", -1},
		{"[-1,0,-2]", 0},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			nums, err := inputs.ReadIntSlice(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			if got := maxSubArray(nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
