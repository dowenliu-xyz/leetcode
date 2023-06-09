package lc0096

import (
	"fmt"
	"testing"
)

func Test_numTrees(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 2},
		{3, 5},
		{19, 1767263190},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			if got := numTrees(tt.n); got != tt.want {
				t.Errorf("numTrees(%v) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
