package main

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "[1,2,3,1]",
			want:  true,
		},
		{
			input: "[1,2,3,4]",
			want:  false,
		},
		{
			input: "[1,1,1,3,3,4,3,2,4,2]",
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			nums := readInput(tt.input)
			got := containsDuplicate(nums)
			if tt.want != got {
				t.Errorf("want %t, got %t", tt.want, got)
			}
		})
	}
}
