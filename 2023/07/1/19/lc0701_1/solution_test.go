package lc0701

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
	"github.com/dowenliu-xyz/leetcode/internal/outputs"
	. "github.com/dowenliu-xyz/leetcode/internal/types"
)

func Test_insertIntoBST(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `
[4,2,7,1,3]
5`,
			want: `[4,2,7,1,3,5]`,
		},
		{
			input: `
[40,20,60,10,30,50,70]
25`,
			want: `[40,20,60,10,30,50,70,null,null,25]`,
		},
		{
			input: `
[4,2,7,1,3,null,null,null,null,null,null]
5`,
			want: `[4,2,7,1,3,5]`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			root, val, err := readInput(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			newRoot := insertIntoBST(root, val)
			got := outputs.LevelOrderString(newRoot)
			if !match(got, tt.want) {
				t.Errorf("insertIntoBST(%v, %v) = %v, want %v", root, val, got, tt.want)
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
		return nil, 0, fmt.Errorf("failed to restore tree: %w", err)
	}
	val, err := strconv.Atoi(strings.TrimSpace(lines[1]))
	if err != nil {
		return nil, 0, fmt.Errorf("failed to parse val: %w", err)
	}
	return root, val, nil
}

func match(s, t string) bool {
	s = strings.ReplaceAll(s, " ", "")
	t = strings.ReplaceAll(t, " ", "")
	return s == t
}
