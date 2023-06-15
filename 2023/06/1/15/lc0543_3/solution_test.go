package lc0543

import (
	"fmt"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
	"github.com/dowenliu-xyz/leetcode/internal/types"
)

func Test_diameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `[1,2,3,4,5]`,
			want:  3,
		},
		{
			input: `[1,2]`,
			want:  1,
		},
		{
			input: `[1,2,null,3,4,5,null,null,6]`,
			want:  4,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			iRoot, err := inputs.LevelOrderRestore(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			root := fromInternal(iRoot)
			if got := diameterOfBinaryTree(root); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func fromInternal(root *types.TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return &TreeNode{
		Val:   root.Val,
		Left:  fromInternal(root.Left),
		Right: fromInternal(root.Right),
	}
}
