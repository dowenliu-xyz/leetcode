package main

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "()",
			want:  true,
		},
		{
			input: "()[]{}",
			want:  true,
		},
		{
			input: "(]",
			want:  false,
		},
		{
			input: "([)]",
			want:  false,
		},
		{
			input: "{[]}",
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := isValid(tt.input)
			if tt.want != got {
				t.Errorf("want %t, got %t", tt.want, got)
			}
		})
	}
}
