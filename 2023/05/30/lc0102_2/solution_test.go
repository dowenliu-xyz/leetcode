package lc0102_2

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
	"github.com/dowenliu-xyz/leetcode/internal/types"
)

func Test_levelOrder(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "[3,9,20,null,null,15,7]",
			want:  "[[3],[9,20],[15,7]]",
		},
		{
			input: "[1]",
			want:  "[[1]]",
		},
		{
			input: "[]",
			want:  "[]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			node, err := inputs.LevelOrderRestore(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			root := fromInternal(node)
			result := levelOrder(root)
			got := formatResult(result)
			if tt.want != got {
				t.Errorf("want %s, got %s", tt.want, got)
			}
		})
	}
}

func fromInternal(node *types.TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	n := &TreeNode{
		Val:   node.Val,
		Left:  fromInternal(node.Left),
		Right: fromInternal(node.Right),
	}
	return n
}

func formatResult(result [][]int) string {
	lines := make([]string, 0, len(result))
	for _, vals := range result {
		line := make([]string, 0, len(vals))
		for _, val := range vals {
			line = append(line, strconv.Itoa(val))
		}
		lines = append(lines, fmt.Sprintf("[%s]", strings.Join(line, ",")))
	}
	return fmt.Sprintf("[%s]", strings.Join(lines, ","))
}
