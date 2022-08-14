package lc0003_1

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "abcabcbb",
			want:  3,
		},
		{
			input: "bbbbb",
			want:  1,
		},
		{
			input: "pwwkew",
			want:  3,
		},
		{
			input: "",
			want:  0,
		},
		{
			input: " ",
			want:  1,
		},
		{
			input: "au",
			want:  2,
		},
		{
			input: "dvdf",
			want:  3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.input)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
