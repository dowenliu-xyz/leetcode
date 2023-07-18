package outputs

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/dowenliu-xyz/leetcode/internal/types"
)

func LevelOrderString(root *TreeNode) string {
	if root == nil {
		return "[]"
	}
	var queue = []*TreeNode{root}
	var res []string
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		res = append(res, strconv.Itoa(node.Val))
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return fmt.Sprintf("[%s]", strings.Join(res, ","))
}
