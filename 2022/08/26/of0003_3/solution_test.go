package main

import "testing"

func TestFindRepeatNumber(t *testing.T) {
	tests := []struct {
		input   string
		wantAny []int
	}{
		{
			input:   "[2, 3, 1, 0, 2, 5, 3]",
			wantAny: []int{2, 3},
		},
		{
			input:   "[3,4,2,1,1,0]",
			wantAny: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			nums := readInput(tt.input)
			got := findRepeatNumber(nums)
			contains := false
			for _, want := range tt.wantAny {
				if want == got {
					contains = true
					break
				}
			}
			if !contains {
				t.Errorf("want any of %v, got %d", tt.wantAny, got)
			}
		})
	}
}
