package lc0111

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
	// bfs
	if root == nil {
		return 0
	}
	ans, q := 0, []*TreeNode{root}
	for {
		ans++
		var p []*TreeNode
		for len(q) > 0 {
			node := q[0]
			q = q[1:]
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
		q = p
	}
}

func solution2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := solution2(root.Left), solution2(root.Right)
	if l == 0 || r == 0 {
		return l + r + 1
	}
	return min(l, r) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
