package main

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "[2,3,1,1,4]",
			want:  true,
		},
		{
			input: "[3,2,1,0,4]",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			nums := readInput(tt.input)
			got := canJump(nums)
			if tt.want != got {
				t.Errorf("want %t, got %t", tt.want, got)
			}
		})
	}
}
