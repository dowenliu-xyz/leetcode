package main

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		nums1 string
		nums2 string
		want  float64
	}{
		{
			nums1: "[1,3]",
			nums2: "[2]",
			want:  2.0,
		},
		{
			nums1: "[1,2]",
			nums2: "[3,4]",
			want:  2.5,
		},
		{
			nums1: "[]",
			nums2: "[]",
			want:  0.0,
		},
		{
			nums1: "[1,2]",
			nums2: "[]",
			want:  1.5,
		},
		{
			nums1: "[]",
			nums2: "[3,4]",
			want:  3.5,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s_%s", tt.nums1, tt.nums2), func(t *testing.T) {
			nums1, nums2 := readInput(tt.nums1), readInput(tt.nums2)
			got := findMedianSortedArrays(nums1, nums2)
			if tt.want != got {
				t.Errorf("want %f, got %f", tt.want, got)
			}
		})
	}
}
