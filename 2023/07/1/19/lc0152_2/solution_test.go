package lc0152

import (
	"fmt"
	"testing"
)

func Test_maxProduct(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if got := maxProduct(tt.nums); got != tt.want {
				t.Errorf("expect %d, got %d", tt.want, got)
			}
		})
	}
}
