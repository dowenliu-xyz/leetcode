package lc0226_5

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
	"github.com/dowenliu-xyz/leetcode/internal/types"
)

func TestInvertTree(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `[4,2,7,1,3,6,9]`,
			want:  `[4,7,2,9,6,3,1]`,
		},
		{
			input: `[2,1,3]`,
			want:  `[2,3,1]`,
		},
		{
			input: `[]`,
			want:  `[]`,
		},
		{
			input: "[4,2,7,null,3,6,9]",
			want:  "[4,7,2,9,6,3]",
		},
		{
			input: "[4,2,7,null,3,6]",
			want:  "[4,7,2,null,6,3]",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			node, err := inputs.LevelOrderRestore(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			root := fromInternal(node)
			invert := invertTree(root)
			got := levelOrderPrint(invert)
			if tt.want != got {
				t.Fatalf("expect %s, got %s", tt.want, got)
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

func levelOrderPrint(node *TreeNode) string {
	if node == nil {
		return "[]"
	}
	var strs []string
	q := []*TreeNode{node}
	for len(q) > 0 {
		var p []*TreeNode
		lastLevel := true
		for _, n := range q {
			if n == nil {
				strs = append(strs, "null")
				continue
			}
			strs = append(strs, strconv.FormatInt(int64(n.Val), 10))
			if n.Left != nil || n.Right != nil {
				lastLevel = false
			}
			p = append(p, n.Left, n.Right)
		}
		if lastLevel {
			break
		}
		q = p
	}
	for ; len(strs) > 0 && strs[len(strs)-1] == "null"; strs = strs[:len(strs)-1] {
	}
	return "[" + strings.Join(strs, ",") + "]"
}
