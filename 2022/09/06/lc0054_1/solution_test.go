package main

import "testing"

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "[[1,2,3],[4,5,6],[7,8,9]]",
			want:  "[1,2,3,6,9,8,7,4,5]",
		},
		{
			input: "[[1,2,3,4],[5,6,7,8],[9,10,11,12]]",
			want:  "[1,2,3,4,8,12,11,10,9,5,6,7]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			matrix := readMatrix(tt.input)
			ints := spiralOrder(matrix)
			got := formatInts(ints)
			if tt.want != got {
				t.Errorf("want %s, got %s", tt.want, got)
			}
		})
	}
}
