package lc0033_2

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func Test_search(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `
[4,5,6,7,0,1,2]
0
`,
			want: 4,
		},
		{
			input: `
[4,5,6,7,0,1,2]
3
`,
			want: -1,
		},
		{
			input: `
[1]
0
`,
			want: -1,
		},
		{
			input: `
[1,3]
0
`,
			want: -1,
		},
		{
			input: `
[3,1]
1
`,
			want: 1,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			nums, target, err := readInput(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			if got := search(nums, target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
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
