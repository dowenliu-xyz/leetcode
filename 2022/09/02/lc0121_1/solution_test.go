package main

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "[7,1,5,3,6,4]",
			want:  5,
		},
		{
			input: "[7,6,4,3,1]",
			want:  0,
		},
		{
			input: "[1]",
			want:  0,
		},
		{
			input: "[]",
			want:  0,
		},
		{
			input: "[1,2]",
			want:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			prices := readInput(tt.input)
			got := maxProfit(prices)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
