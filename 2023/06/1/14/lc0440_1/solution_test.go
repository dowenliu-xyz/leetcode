package lc0440

import (
	"fmt"
	"testing"
)

func Test_findKthNumber(t *testing.T) {
	tests := []struct {
		n    int
		k    int
		want int
	}{
		{13, 2, 10},
		{13, 3, 11},
		{1, 1, 1},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if got := findKthNumber(tt.n, tt.k); got != tt.want {
				t.Errorf("findKthNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
