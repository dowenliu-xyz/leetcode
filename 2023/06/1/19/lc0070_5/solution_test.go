package lc0070

import (
	"fmt"
	"testing"
)

func Test_climbStairs(t *testing.T) {
	tests := []struct {
		input  int
		expect int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if got := climbStairs(tt.input); got != tt.expect {
				t.Errorf("climbStairs() = %v, want %v", got, tt.expect)
			}
		})
	}
}
