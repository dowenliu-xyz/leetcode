package main

import "testing"

func TestMerge(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "[[1,3],[2,6],[8,10],[15,18]]",
			want:  "[[1,6],[8,10],[15,18]]",
		},
		{
			input: "[[1,4],[4,5]]",
			want:  "[[1,5]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			intervals := readIntervals(tt.input)
			intervals = merge(intervals)
			got := formatIntervals(intervals)
			if tt.want != got {
				t.Errorf("want %s, got %s", tt.want, got)
			}
		})
	}
}
