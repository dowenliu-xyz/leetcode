package lc0104_2

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
	if root == nil {
		return 0
	}
	ans, q := 0, []*TreeNode{root}
	for len(q) > 0 {
		ans++
		var p []*TreeNode
		for len(q) > 0 {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ans
}

func solution2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := solution2(root.Left), solution2(root.Right)
	return max(l, r) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
