package lc0215_1

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func Test_findKthLargest(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `
[3,2,1,5,6,4]
2
`,
			want: 5,
		},
		{
			input: `
[3,2,3,1,2,4,5,5,6]
4
`,
			want: 4,
		},
		{
			input: `
[1]
1
`,
			want: 1,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			nums, k, err := readInput(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			got := findKthLargest(nums, k)
			if got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
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
