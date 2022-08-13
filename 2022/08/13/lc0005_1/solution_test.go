package main

import (
	"strings"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  map[string]bool
	}{
		{
			input: "babad",
			want: map[string]bool{
				"bab": true,
				"aba": true,
			},
		},
		{
			input: "cbbd",
			want: map[string]bool{
				"bb": true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := longestPalindrome(tt.input)
			if _, ok := tt.want[got]; !ok {
				want := make([]string, 0, len(tt.want))
				for s := range tt.want {
					want = append(want, s)
				}
				t.Errorf("want any of [%s], got %s", strings.Join(want, ","), got)
			}
		})
	}
}
