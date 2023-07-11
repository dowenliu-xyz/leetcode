package lc0016

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func Test_threeSumClosest(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `
[-1,2,1,-4]
1
`,
			want: 2,
		},
		{
			input: `
[0,0,0]
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
			if got := threeSumClosest(nums, target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func readInput(input string) (nums []int, target int, err error) {
	lines := inputs.ReadNoneBlankLines(input)
	if len(lines) != 2 {
		err = fmt.Errorf("input should have 2 lines")
		return
	}
	nums, err = inputs.ReadIntSlice(lines[0])
	if err != nil {
		return
	}
	target, err = strconv.Atoi(strings.TrimSpace(lines[1]))
	if err != nil {
		return
	}
	return
}
