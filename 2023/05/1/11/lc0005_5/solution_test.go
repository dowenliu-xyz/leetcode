package lc0005_5

import (
	"fmt"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `babad`,
			want:  `bab`,
		},
		{
			input: `cbbd`,
			want:  `bb`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := longestPalindrome(tt.input)
			if got != tt.want {
				t.Errorf("expect %s, got %s", tt.want, got)
			}
		})
	}
}
