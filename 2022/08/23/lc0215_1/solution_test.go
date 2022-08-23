package main

import (
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		nums string
		k    int
		want int
	}{
		{
			nums: `[3,2,1,5,6,4]`,
			k:    2,
			want: 5,
		},
		{
			nums: `[3,2,3,1,2,4,5,5,6]`,
			k:    4,
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s_%d", tt.nums, tt.k), func(t *testing.T) {
			nums := readInput(tt.nums)
			got := findKthLargest(nums, tt.k)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
