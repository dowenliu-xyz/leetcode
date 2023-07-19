package outputs

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/samber/lo"

	. "github.com/dowenliu-xyz/leetcode/internal/types"
)

func LevelOrderString(root *TreeNode) string {
	var queue = []*TreeNode{root}
	var res []string
	for len(queue) > 0 {
		var newQueue []*TreeNode
		var level []string
		for _, node := range queue {
			if node == nil {
				level = append(level, "null")
			} else {
				level = append(level, strconv.Itoa(node.Val))
				newQueue = append(newQueue, node.Left, node.Right)
			}
		}
		if lo.EveryBy(level, func(s string) bool { return s == "null" }) {
			queue = nil
			continue
		}
		queue = newQueue
		res = append(res, level...)
	}
	for len(res) > 0 && res[len(res)-1] == "null" {
		res = res[:len(res)-1]
	}
	return fmt.Sprintf("[%s]", strings.Join(res, ","))
}
