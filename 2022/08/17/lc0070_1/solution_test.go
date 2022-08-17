package main

import (
	"fmt"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{
			input: 0,
			want:  1,
		},
		{
			input: 1,
			want:  1,
		},
		{
			input: 2,
			want:  2,
		},
		{
			input: 3,
			want:  3,
		},
		{
			input: 4,
			want:  5,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.input), func(t *testing.T) {
			got := climbStairs(tt.input)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
