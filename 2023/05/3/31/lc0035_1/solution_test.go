package lc0035_1

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func Test_searchInsert(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `
[1,3,5,6]
5
`,
			want: 2,
		},
		{
			input: `
[1,3,5,6]
2
`,
			want: 1,
		},
		{
			input: `
[1,3,5,6]
7
`,
			want: 4,
		},
		{
			input: `
[1,3,5,6]
0
`,
			want: 0,
		},
		{
			input: `
[]
1
`,
			want: 0,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			nums, target, err := readInput(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			if got := searchInsert(nums, target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func readInput(input string) ([]int, int, error) {
	lines := inputs.ReadNoneBlankLines(input)
	if len(lines) != 2 {
		return nil, 0, fmt.Errorf("input should have 2 lines")
	}
	nums, err := inputs.ReadIntSlice(lines[0])
	if err != nil {
		return nil, 0, err
	}
	target, err := strconv.Atoi(strings.TrimSpace(lines[1]))
	if err != nil {
		return nil, 0, err
	}
	return nums, target, nil
}
