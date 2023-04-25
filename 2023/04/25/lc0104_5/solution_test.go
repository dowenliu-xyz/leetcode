package lc0104_5

import (
	"fmt"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
	"github.com/dowenliu-xyz/leetcode/internal/types"
)

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `[3,9,20,null,null,15,7]`,
			want:  3,
		},
		{
			input: "[]",
			want:  0,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			node, err := inputs.LevelOrderRestore(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			root := fromInternal(node)
			if got := maxDepth(root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
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
