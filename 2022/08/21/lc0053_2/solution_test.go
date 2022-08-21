package main

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "[-2,1,-3,4,-1,2,1,-5,4]",
			want:  6,
		},
		{
			input: "[1]",
			want:  1,
		},
		{
			input: "[5,4,-1,7,8]",
			want:  23,
		},
		{
			input: "[-1]",
			want:  -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			nums := readInput(tt.input)
			got := maxSubArray(nums)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
