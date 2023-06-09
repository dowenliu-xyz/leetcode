package lc0560

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func Test_subarraySum(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `
[1,1,1]
2
`,
			want: 2,
		},
		{
			input: `
[1,2,3]
3
`,
			want: 2,
		},
		{
			input: `
[1,-1,0]
0
`,
			want: 3,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			nums, k, err := readInput(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			if got := subarraySum(nums, k); got != tt.want {
				t.Errorf("subarraySum() = %v, want %v", got, tt.want)
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
	k, err := strconv.Atoi(strings.TrimSpace(lines[1]))
	if err != nil {
		return nil, 0, err
	}
	return nums, k, nil
}
