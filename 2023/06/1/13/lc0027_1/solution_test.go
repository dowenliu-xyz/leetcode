package lc0027

import (
	"fmt"
	"sort"
	"testing"
)

func Test_removeElement(t *testing.T) {
	tests := []struct {
		nums     []int
		val      int
		wantLen  int
		wantNums []int
	}{
		{
			nums:     []int{3, 2, 2, 3},
			val:      3,
			wantLen:  2,
			wantNums: []int{2, 2},
		},
		{
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			wantLen:  5,
			wantNums: []int{0, 1, 3, 0, 4},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			gotLen := removeElement(tt.nums, tt.val)
			if gotLen != tt.wantLen {
				t.Errorf("removeElement(%v, %d) = %v, want %v", tt.nums, tt.val, gotLen, tt.wantLen)
			}
			sort.Ints(tt.nums[:gotLen])
			sort.Ints(tt.wantNums)
			for i := 0; i < tt.wantLen; i++ {
				if tt.nums[i] != tt.wantNums[i] {
					t.Errorf("removeElement(%v, %d) = %v, want %v", tt.nums, tt.val, tt.nums, tt.wantNums)
				}
			}
		})
	}
}
