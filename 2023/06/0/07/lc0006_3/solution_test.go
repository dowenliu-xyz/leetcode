package lc0006

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func Test_convert(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `
PAYPALISHIRING
3
`,
			want: "PAHNAPLSIIGYIR",
		},
		{
			input: `
PAYPALISHIRING
4
`,
			want: "PINALSIGYAHRPI",
		},
		{
			input: `
A
1
`,
			want: "A",
		},
		{
			input: `
AB
1
`,
			want: "AB",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			s, numRows, err := readInput(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			got := convert(s, numRows)
			if got != tt.want {
				t.Fatalf("want: %s, got: %s", tt.want, got)
			}
		})
	}
}

func readInput(input string) (string, int, error) {
	lines := inputs.ReadNoneBlankLines(input)
	if len(lines) != 2 {
		return "", 0, fmt.Errorf("input should have 2 lines")
	}
	num, err := strconv.Atoi(strings.TrimSpace(lines[1]))
	if err != nil {
		return "", 0, fmt.Errorf("invalid number of rows: %w", err)
	}
	return strings.TrimSpace(lines[0]), num, nil
}
