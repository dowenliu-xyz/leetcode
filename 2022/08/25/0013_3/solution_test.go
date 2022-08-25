package main

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "III",
			want:  3,
		},
		{
			input: "IV",
			want:  4,
		},
		{
			input: "IX",
			want:  9,
		},
		{
			input: "LVIII",
			want:  58,
		},
		{
			input: "MCMXCIV",
			want:  1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := romanToInt(tt.input)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
