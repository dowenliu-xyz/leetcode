package lc0100_2

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	return solution2(p, q)
}

func solution1(p, q *TreeNode) bool {
	// 递归，dfs
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return solution1(p.Left, q.Left) && solution1(p.Right, q.Right)
}

func solution2(p, q *TreeNode) bool {
	// 迭代，bfs
	queue1, queue2 := []*TreeNode{p}, []*TreeNode{q}
	for len(queue1) > 0 && len(queue2) > 0 {
		node1, node2 := queue1[0], queue2[0]
		queue1, queue2 = queue1[1:], queue2[1:]
		if (node1 == nil && node2 != nil) || (node1 != nil && node2 == nil) {
			return false
		}
		if (node1 != nil && node2 != nil) && node1.Val != node2.Val {
			return false
		}
		if node1 != nil {
			queue1 = append(queue1, node1.Left, node1.Right)
		}
		if node2 != nil {
			queue2 = append(queue2, node2.Left, node2.Right)
		}
	}
	return len(queue1) == 0 && len(queue2) == 0
}
