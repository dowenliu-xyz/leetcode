package main

import "testing"

func TestRunningSum(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `[1,2,3,4]`,
			want:  `[1,3,6,10]`,
		},
		{
			input: `[1,1,1,1,1]`,
			want:  `[1,2,3,4,5]`,
		},
		{
			input: `[3,1,2,10,1]`,
			want:  `[3,4,6,16,17]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			nums := readInput(tt.input)
			sum := runningSum(nums)
			got := formatInts(sum)
			if tt.want != got {
				t.Errorf("want %s, got %s", tt.want, got)
			}
		})
	}
}
