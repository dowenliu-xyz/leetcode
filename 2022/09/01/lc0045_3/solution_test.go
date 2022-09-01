package main

import "testing"

func TestJump(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `[2,3,1,1,4]`,
			want:  2,
		},
		{
			input: `[2,3,0,1,4]`,
			want:  2,
		},
		{
			input: `[0]`,
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			nums := readInput(tt.input)
			got := jump(nums)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
