package lc0704_5

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
[-1,0,3,5,9,12]
9
`,
			want: 4,
		},
		{
			input: `
[-1,0,3,5,9,12]
2
`,
			want: -1,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			nums, target, err := readInput(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			got := search(nums, target)
			if got != tt.want {
				t.Fatalf("expect %d, got %d", tt.want, got)
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
