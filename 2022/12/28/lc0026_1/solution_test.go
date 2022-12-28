package main

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input    string
		wantNums []int
		wantLen  int
	}{
		{
			input:    `[1,1,2]`,
			wantLen:  2,
			wantNums: []int{1, 2},
		},
		{
			input:    `[0,0,1,1,1,2,2,3,3,4]`,
			wantLen:  5,
			wantNums: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			nums := readInput(tt.input)
			l := removeDuplicates(nums)
			nums = nums[:l]
			if tt.wantLen != l {
				t.Errorf("expect len %d, got %d", tt.wantLen, l)
			}
			if !reflect.DeepEqual(nums, tt.wantNums) {
				t.Errorf("expect %v, got %v", tt.wantNums, nums)
			}
		})
	}
}
