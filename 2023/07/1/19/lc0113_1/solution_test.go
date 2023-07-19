package lc0113

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
	"github.com/dowenliu-xyz/leetcode/internal/outputs"
	. "github.com/dowenliu-xyz/leetcode/internal/types"
)

func Test_pathSum(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `
[5,4,8,11,null,13,4,7,2,null,null,5,1]
22
`,
			want: `[[5,4,11,2],[5,8,4,5]]`,
		},
		{
			input: `
[1,2,3]
5
`,
			want: `[]`,
		},
		{
			input: `
[1,2]
0
`,
			want: `[]`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			root, targetSum, err := readInput(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			paths := pathSum(root, targetSum)
			got := outputs.SprintIntMatrix(paths)
			if !match(got, tt.want) {
				t.Errorf("pathSum(%v, %v) = %v, want %v", root, targetSum, got, tt.want)
			}
		})
	}
}

func readInput(input string) (*TreeNode, int, error) {
	lines := inputs.ReadNoneBlankLines(input)
	if len(lines) != 2 {
		return nil, 0, fmt.Errorf("input should have 2 lines")
	}
	root, err := inputs.LevelOrderRestore(lines[0])
	if err != nil {
		return nil, 0, fmt.Errorf("read tree root error: %w", err)
	}
	targetSum, err := strconv.Atoi(strings.TrimSpace(lines[1]))
	if err != nil {
		return nil, 0, fmt.Errorf("read target sum error: %w", err)
	}
	return root, targetSum, nil
}

func match(got, want string) bool {
	got = strings.ReplaceAll(got, " ", "")
	want = strings.ReplaceAll(want, " ", "")
	return got == want
}
