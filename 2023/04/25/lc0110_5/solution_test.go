package lc0110_5

import (
	"fmt"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
	"github.com/dowenliu-xyz/leetcode/internal/types"
)

func TestIsBalance(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: `[3,9,20,null,null,15,7]`,
			want:  true,
		},
		{
			input: `[1,2,2,3,3,null,null,4,4]`,
			want:  false,
		},
		{
			input: `[]`,
			want:  true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			node, err := inputs.LevelOrderRestore(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			root := fromInternal(node)
			got := isBalanced(root)
			if tt.want != got {
				t.Errorf("expect %t, got %t", tt.want, got)
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
