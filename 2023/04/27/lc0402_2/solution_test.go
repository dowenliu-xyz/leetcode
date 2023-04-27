package lc0402_2

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func parseInput(t *testing.T, input string) (string, int) {
	t.Helper()

	lines := inputs.ReadNoneBlankLines(input)
	if len(lines) != 2 {
		t.Fatalf("expect 2 lines, got %d", len(lines))
	}
	num, err := strconv.Unquote(lines[0])
	if err != nil {
		t.Fatal(err)
	}
	k, err := strconv.ParseInt(lines[1], 10, 64)
	if err != nil {
		t.Fatal(err)
	}
	return num, int(k)
}

func TestRemoveKdigits(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `
"1432219"
3
`,
			want: "1219",
		},
		{
			input: `
"10200"
1
`,
			want: "200",
		},
		{
			input: `
"10"
2
`,
			want: "0",
		},
		{
			input: `
"112"
1
`,
			want: "11",
		},
		{
			input: `
"1234567890"
9
`,
			want: "0",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			num, k := parseInput(t, tt.input)
			got := removeKdigits(num, k)
			if got != tt.want {
				t.Errorf("tests[%d] got %q, want %q", i, got, tt.want)
			}
		})
	}
}
