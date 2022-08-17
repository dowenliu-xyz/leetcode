package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	//return solution1(root)
	return solution2(root)
}

func solution1(root *TreeNode) int {
	// dfs
	if root == nil {
		return 0
	}
	if root.Left == nil || root.Right == nil {
		return maxDepth(root.Left) + maxDepth(root.Right) + 1
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func solution2(root *TreeNode) int {
	// bfs
	if root == nil {
		return 0
	}
	ans, q := 0, []*TreeNode{root}
	for len(q) > 0 {
		var p []*TreeNode
		for _, node := range q {
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		ans++
		q = p
	}
	return ans
}
