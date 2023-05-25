package lc0004_2

import (
	"fmt"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func Test_findMedianSortedArrays(t *testing.T) {
	tests := []struct {
		input string
		want  float64
	}{
		{
			input: `
[1,3]
[2]
`,
			want: 2.0,
		},
		{
			input: `
[1,2]
[3,4]
`,
			want: 2.5,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			nums1, nums2, err := readInput(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			got := findMedianSortedArrays(nums1, nums2)
			if got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func readInput(input string) ([]int, []int, error) {
	lines := inputs.ReadNoneBlankLines(input)
	if len(lines) != 2 {
		return nil, nil, fmt.Errorf("input should have 2 lines")
	}
	nums1, err := inputs.ReadIntSlice(lines[0])
	if err != nil {
		return nil, nil, err
	}
	nums2, err := inputs.ReadIntSlice(lines[1])
	if err != nil {
		return nil, nil, err
	}
	return nums1, nums2, nil
}
