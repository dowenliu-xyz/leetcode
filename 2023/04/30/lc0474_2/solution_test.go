package lc0474_2

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func TestFindMaxForm(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `
["10","0001","111001","1","0"]
5
3
`,
			want: 4,
		},
		{
			input: `
["10","0","1"]
1
1
`,
			want: 2,
		},
		{
			input: `
["00","000"]
1
10
`,
			want: 0,
		},
		{
			input: `
["10","0001","111001","1","0"]
3
4
`,
			want: 3,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			strs, m, n, err := parseInput(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			got := findMaxForm(strs, m, n)
			if got != tt.want {
				t.Fatalf("want: %d, got: %d", tt.want, got)
			}
		})
	}
}

func parseInput(input string) (strs []string, m, n int, err error) {
	lines := inputs.ReadNoneBlankLines(input)
	if len(lines) != 3 {
		err = fmt.Errorf("input should have 3 lines")
		return
	}
	strs, err = inputs.ReadStringSlice(lines[0])
	if err != nil {
		return
	}
	m, err = strconv.Atoi(lines[1])
	if err != nil {
		return
	}
	n, err = strconv.Atoi(lines[2])
	return
}
