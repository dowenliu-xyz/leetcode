package lc0278_1

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func Test_firstBadVersion(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `
5
4
`,
			want: 4,
		},
		{
			input: `
1
1
`,
			want: 1,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			n, f, err := readInput(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			isBadVersion = f
			if got := firstBadVersion(n); got != tt.want {
				t.Fatalf("want: %d, got: %d", tt.want, got)
			}
		})
	}
}

func readInput(input string) (int, func(int) bool, error) {
	lines := inputs.ReadNoneBlankLines(input)
	if len(lines) != 2 {
		return 0, nil, fmt.Errorf("input should have 2 lines")
	}
	n, err := strconv.Atoi(strings.TrimSpace(lines[0]))
	if err != nil {
		return 0, nil, err
	}
	v, err := strconv.Atoi(strings.TrimSpace(lines[1]))
	if err != nil {
		return 0, nil, err
	}
	return n, func(i int) bool {
		return i >= v
	}, nil
}
