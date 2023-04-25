package lc0100_5

import (
	"fmt"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
	"github.com/dowenliu-xyz/leetcode/internal/types"
)

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: `
[1,2,3]
[1,2,3]
`,
			want: true,
		},
		{
			input: `
[1,2]
[1,null,2]
`,
			want: false,
		},
		{
			input: `
[1,2,1]
[1,1,2]
`,
			want: false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			lines := inputs.ReadNoneBlankLines(tt.input)
			if len(lines) != 2 {
				t.Fatalf("invalid input")
			}
			node1, err := inputs.LevelOrderRestore(lines[0])
			if err != nil {
				t.Fatalf("invalid input: %v", err)
			}
			node2, err := inputs.LevelOrderRestore(lines[1])
			if err != nil {
				t.Fatalf("invalid input: %v", err)
			}
			p, q := fromInternal(node1), fromInternal(node2)
			got := isSameTree(p, q)
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
