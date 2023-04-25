package lc0101_5

import (
	"fmt"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
	"github.com/dowenliu-xyz/leetcode/internal/types"
)

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: `[1,2,2,3,4,4,3]`,
			want:  true,
		},
		{
			input: `[1,2,2,null,3,null,3]`,
			want:  false,
		},
		{
			input: "[]",
			want:  true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			node, err := inputs.LevelOrderRestore(tt.input)
			if err != nil {
				t.Fatalf("invalid input: %v", err)
			}
			root := fromInternal(node)
			got := isSymmetric(root)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
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
