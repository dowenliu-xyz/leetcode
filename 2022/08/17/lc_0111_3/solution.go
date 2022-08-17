package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	//return solution1(root)
	return solution2(root)
}

func solution1(root *TreeNode) int {
	// dfs
	if root == nil {
		return 0
	}
	if root.Left == nil || root.Right == nil {
		return minDepth(root.Left) + minDepth(root.Right) + 1
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func solution2(root *TreeNode) int {
	// bfs
	if root == nil {
		return 0
	}
	ans, q := 1, []*TreeNode{root}
	for len(q) > 0 {
		var p []*TreeNode
		for _, node := range q {
			if node.Left == nil && node.Right == nil {
				return ans
			}
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
