package lc0104

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
	l, r := solution1(root.Left), solution1(root.Right)
	return max(l, r) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solution2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		ans++
		var p []*TreeNode
		for _, node := range q {
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
