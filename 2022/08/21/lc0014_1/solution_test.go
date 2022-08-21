package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `["flower","flow","flight"]`,
			want:  "fl",
		},
		{
			input: `["dog","racecar","car"]`,
			want:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			strs := parseStrings(tt.input)
			got := longestCommonPrefix(strs)
			if tt.want != got {
				t.Errorf("want %s, got %s", tt.want, got)
			}
		})
	}
}
