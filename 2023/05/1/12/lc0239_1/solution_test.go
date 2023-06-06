package lc0239_1

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
	"github.com/dowenliu-xyz/leetcode/internal/outputs"
)

func Test_maxSlidingWindow(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `
[1,3,-1,-3,5,3,6,7]
3
`,
			want: `[3,3,5,5,6,7]`,
		},
		{
			input: `
[1]
1
`,
			want: `[1]`,
		},
		{
			input: `
[1,3,1,2,0,5]
3
`,
			want: `[3,3,2,5]`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			nums, k, err := readInput(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			got := maxSlidingWindow(nums, k)
			want := strings.ReplaceAll(tt.want, " ", "")
			if got := outputs.SprintIntSlice(got); got != want {
				t.Errorf("expect %s, got %s", want, got)
			}
		})
	}
}

func readInput(input string) ([]int, int, error) {
	lines := inputs.ReadNoneBlankLines(input)
	if len(lines) != 2 {
		return nil, 0, fmt.Errorf("input should have 2 lines, but got %d", len(lines))
	}
	nums, err := inputs.ReadIntSlice(lines[0])
	if err != nil {
		return nil, 0, err
	}
	k, err := strconv.Atoi(strings.TrimSpace(lines[1]))
	if err != nil {
		return nil, 0, err
	}
	return nums, k, nil
}
