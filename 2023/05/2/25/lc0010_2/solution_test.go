package lc0010_2

import (
	"fmt"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func TestIsMatch(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: `
aa
a
`,
			want: false,
		},
		{
			input: `
aaa
a*
`,
			want: true,
		},
		{
			input: `
ab
.*
`,
			want: true,
		},
		{
			input: `
mississippi
mis*is*ip*.
`,
			want: true,
		},
		{
			input: `
"mississippi"
"mis*is*p*."
`,
			want: false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			lines := inputs.ReadNoneBlankLines(tt.input)
			if len(lines) != 2 {
				t.Fatal("input should have 2 lines")
			}
			s := lines[0]
			p := lines[1]
			if got := isMatch(s, p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
