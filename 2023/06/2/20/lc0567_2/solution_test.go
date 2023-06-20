package lc0567

import (
	"fmt"
	"testing"
)

func Test_checkInclusion(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want bool
	}{
		{"ab", "eidbaooo", true},
		{"ab", "eidboaoo", false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if got := checkInclusion(tt.s1, tt.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}
