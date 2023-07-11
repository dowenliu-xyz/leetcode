package lc0069

import (
	"fmt"
	"testing"
)

func Test_mySqrt(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{4, 2},
		{8, 2},
		{0, 0},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if got := mySqrt(tt.x); got != tt.want {
				t.Errorf("mySqrt(%d) = %v, want %v", tt.x, got, tt.want)
			}
		})
	}
}
