package lc0704_1

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   string
		target int
		want   int
	}{
		{
			nums:   "[-1,0,3,5,9,12]",
			target: 9,
			want:   4,
		},
		{
			nums:   "[-1,0,3,5,9,12]",
			target: 2,
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s_%d", tt.nums, tt.target), func(t *testing.T) {
			nums := readInput(tt.nums)
			got := search(nums, tt.target)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
