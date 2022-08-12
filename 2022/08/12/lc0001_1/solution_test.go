package lc0001_1

import (
	"fmt"
	"strconv"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   string
		target string
		want   string
	}{
		{
			nums:   "[2,7,11,15]",
			target: "9",
			want:   "[0,1]",
		},
		{
			nums:   "[3,2,4]",
			target: "6",
			want:   "[1,2]",
		},
		{
			nums:   "[3,3]",
			target: "6",
			want:   "[0,1]",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s_%s", tt.nums, tt.target), func(t *testing.T) {
			nums := readInput(tt.nums)
			target, err := strconv.Atoi(tt.target)
			if err != nil {
				panic(err)
			}
			result := twoSum(nums, target)
			got := formatInts(result)
			if tt.want != got {
				t.Errorf("want %s, got %s", tt.want, got)
			}
		})
	}
}
