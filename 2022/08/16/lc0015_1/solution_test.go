package main

import (
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "[-1,0,1,2,-1,-4]",
			want:  "[[-1,-1,2],[-1,0,1]]",
		},
		{
			input: "[]",
			want:  "[]",
		},
		{
			input: "[0]",
			want:  "[]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			nums := readInput(tt.input)
			result := threeSum(nums)
			sort.Slice(result, func(i, j int) bool {
				for k := 0; k < 3; k++ {
					if result[i][k] > result[j][k] {
						return false
					}
				}
				return true
			})
			got := formatResult(result)
			if tt.want != got {
				t.Errorf("want %s, got %s", tt.want, got)
			}
		})
	}
}
