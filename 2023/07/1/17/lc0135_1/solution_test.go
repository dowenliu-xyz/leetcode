package lc0135

import (
	"fmt"
	"testing"
)

func Test_candy(t *testing.T) {
	tests := []struct {
		ratings []int
		want    int
	}{
		{[]int{1, 0, 2}, 5},
		{[]int{1, 2, 2}, 4},
		{[]int{1, 3, 5, 2, 4, 6}, 12},
		{[]int{1, 3, 5, 4, 2, 6}, 11},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if got := candy(tt.ratings); got != tt.want {
				t.Errorf("expect %d, got %d", tt.want, got)
			}
		})
	}
}
