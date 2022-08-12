package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	//return solution1(root)
	return solution2(root)
}

func solution1(root *TreeNode) bool {
	// 递归。dfs
	return isMirror(root, root)
}

func isMirror(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return isMirror(root1.Left, root2.Right) && isMirror(root1.Right, root2.Left)
}

func solution2(root *TreeNode) bool {
	// 迭代 bfs
	q := []*TreeNode{root, root}
	for len(q) > 1 {
		node1, node2 := q[0], q[1]
		q = q[2:]
		if node1 == nil && node2 == nil {
			continue
		}
		if node1 == nil || node2 == nil {
			return false
		}
		if node1.Val != node2.Val {
			return false
		}
		q = append(q, node1.Left, node2.Right, node1.Right, node2.Left)
	}
	return len(q) == 0
}
