package lc0026_3

import (
	"fmt"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type want struct {
		length int
		nums   []int
	}
	tests := []struct {
		nums []int
		want want
	}{
		{
			nums: []int{1, 1, 2},
			want: want{
				length: 2,
				nums:   []int{1, 2},
			},
		},
		{
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: want{
				length: 5,
				nums:   []int{0, 1, 2, 3, 4},
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			length := removeDuplicates(tt.nums)
			if length != tt.want.length {
				t.Errorf("got length = %v, want %v", length, tt.want.length)
			}
			for i := 0; i < length; i++ {
				if tt.nums[i] != tt.want.nums[i] {
					t.Errorf("got nums = %v, want %v", tt.nums[:length], tt.want.nums)
					break
				}
			}
		})
	}
}
