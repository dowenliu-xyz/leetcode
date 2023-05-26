package lc0014_3

import (
	"fmt"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func Test_longestCommonPrefix(t *testing.T) {
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
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			strs, err := inputs.ReadStringSlice(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			if got := longestCommonPrefix(strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
