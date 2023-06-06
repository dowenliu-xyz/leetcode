package lc0088_2

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
	"github.com/dowenliu-xyz/leetcode/internal/outputs"
)

func Test_merge(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `
[1,2,3,0,0,0]
3
[2,5,6]
3
`,
			want: "[1,2,2,3,5,6]",
		},
		{
			input: `
[1]
1
[]
0
`,
			want: "[1]",
		},
		{
			input: `
[0]
0
[1]
1
`,
			want: "[1]",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			nums1, m, nums2, n, err := readInput(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			merge(nums1, m, nums2, n)
			if got := outputs.SprintIntSlice(nums1); got != tt.want {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func readInput(input string) ([]int, int, []int, int, error) {
	lines := inputs.ReadNoneBlankLines(input)
	if len(lines) != 4 {
		return nil, 0, nil, 0, fmt.Errorf("input should have 4 lines")
	}
	nums1, err := inputs.ReadIntSlice(lines[0])
	if err != nil {
		return nil, 0, nil, 0, err
	}
	m, err := strconv.Atoi(strings.TrimSpace(lines[1]))
	if err != nil {
		return nil, 0, nil, 0, err
	}
	nums2, err := inputs.ReadIntSlice(lines[2])
	if err != nil {
		return nil, 0, nil, 0, err
	}
	n, err := strconv.Atoi(strings.TrimSpace(lines[3]))
	if err != nil {
		return nil, 0, nil, 0, err
	}
	if len(nums2) != n {
		return nil, 0, nil, 0, fmt.Errorf("nums2 should have %d elements", n)
	}
	if len(nums1) != m+n {
		return nil, 0, nil, 0, fmt.Errorf("nums1 should have %d elements", m+n)
	}
	return nums1, m, nums2, n, nil
}
