package lc0236

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
	"github.com/dowenliu-xyz/leetcode/internal/types"
)

func Test_lowestCommonAncestor(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `
[3,5,1,6,2,0,8,null,null,7,4]
5
1
`,
			want: 3,
		},
		{
			input: `
[3,5,1,6,2,0,8,null,null,7,4]
5
4
`,
			want: 5,
		},
		{
			input: `
[1,2]
1
2
`,
			want: 1,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			root, p, q, err := readInput(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			node := lowestCommonAncestor(root, p, q)
			if node == nil {
				t.Fatal("node should not be nil")
			}
			if node.Val != tt.want {
				t.Errorf("expect %d, got %d", tt.want, node.Val)
			}
		})
	}
}

func readInput(input string) (*TreeNode, *TreeNode, *TreeNode, error) {
	lines := inputs.ReadNoneBlankLines(input)
	if len(lines) != 3 {
		return nil, nil, nil, fmt.Errorf("input should have 3 lines, but got %d", len(lines))
	}
	internalRoot, err := inputs.LevelOrderRestore(lines[0])
	if err != nil {
		return nil, nil, nil, err
	}
	root := fromInternal(internalRoot)
	pv, err := strconv.Atoi(lines[1])
	if err != nil {
		return nil, nil, nil, err
	}
	p := findOfV(root, pv)
	if p == nil {
		return nil, nil, nil, fmt.Errorf("can not find node with value %d", pv)
	}
	qv, err := strconv.Atoi(lines[2])
	if err != nil {
		return nil, nil, nil, err
	}
	q := findOfV(root, qv)
	if q == nil {
		return nil, nil, nil, fmt.Errorf("can not find node with value %d", qv)
	}
	return root, p, q, nil
}

func findOfV(root *TreeNode, v int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == v {
		return root
	}
	l := findOfV(root.Left, v)
	if l != nil {
		return l
	}
	r := findOfV(root.Right, v)
	if r != nil {
		return r
	}
	return nil
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
