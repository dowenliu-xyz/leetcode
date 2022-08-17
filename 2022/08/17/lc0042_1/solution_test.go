package main

import "testing"

func TestTrap(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "[0,1,0,2,1,0,1,3,2,1,2,1]",
			want:  6,
		},
		{
			input: "[4,2,0,3,2,5]",
			want:  9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			height := readInput(tt.input)
			got := trap(height)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
