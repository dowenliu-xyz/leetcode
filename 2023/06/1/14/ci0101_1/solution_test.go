package ci0101

import (
	"fmt"
	"testing"
)

func Test_isUnique(t *testing.T) {
	tests := []struct {
		astr string
		want bool
	}{
		{"leetcode", false},
		{"abc", true},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if got := isUnique(tt.astr); got != tt.want {
				t.Errorf("isUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
