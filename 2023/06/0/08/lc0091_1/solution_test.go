package lc0091

import "testing"

func Test_numTrees(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 2},
		{3, 5},
	}
	for _, tt := range tests {
		if got := numTrees(tt.n); got != tt.want {
			t.Errorf("numTrees(%v) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
