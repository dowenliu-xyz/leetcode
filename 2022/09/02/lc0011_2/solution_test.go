package main

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "[1,8,6,2,5,4,8,3,7]",
			want:  49,
		},
		{
			input: "[1,1]",
			want:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			height := readInput(tt.input)
			got := maxArea(height)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
