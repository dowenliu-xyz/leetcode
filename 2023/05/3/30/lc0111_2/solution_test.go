package lc0111

import (
	"fmt"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
	"github.com/dowenliu-xyz/leetcode/internal/types"
)

func TestMinDepth(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `[3,9,20,null,null,15,7]`,
			want:  2,
		},
		{
			input: `[2,null,3,null,4,null,5,null,6]`,
			want:  5,
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
				t.Fatalf("unexpect err: %+v", err)
			}
			root := fromInternal(node)
			got := minDepth(root)
			if got != tt.want {
				t.Fatalf("expect %d, got %d", tt.want, got)
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
