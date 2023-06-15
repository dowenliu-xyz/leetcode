package lc0124

import (
	"fmt"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
	"github.com/dowenliu-xyz/leetcode/internal/types"
)

func Test_maxPathSum(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "[1,2,3]",
			want:  6,
		},
		{
			input: "[-10,9,20,null,null,15,7]",
			want:  42,
		},
		{
			input: "[-3]",
			want:  -3,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			node, err := inputs.LevelOrderRestore(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			got := maxPathSum(fromInternal(node))
			if got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func fromInternal(node *types.TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	return &TreeNode{
		Val:   node.Val,
		Left:  fromInternal(node.Left),
		Right: fromInternal(node.Right),
	}
}

func toInternal(root *TreeNode) *types.TreeNode {
	if root == nil {
		return nil
	}
	return &types.TreeNode{
		Val:   root.Val,
		Left:  toInternal(root.Left),
		Right: toInternal(root.Right),
	}
}
