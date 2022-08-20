package main

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{
			input: 121,
			want:  true,
		},
		{
			input: -121,
			want:  false,
		},
		{
			input: 10,
			want:  false,
		},
		{
			input: 0,
			want:  true,
		},
		{
			input: 1001,
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.input), func(t *testing.T) {
			got := isPalindrome(tt.input)
			if tt.want != got {
				t.Errorf("want %t, got %t", tt.want, got)
			}
		})
	}
}
