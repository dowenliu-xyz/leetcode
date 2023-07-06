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
	//return dfs(root)
	return bfs(root)
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := dfs(root.Left), dfs(root.Right)
	return max(l, r) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func bfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		var p []*TreeNode
		ans++
		for _, root := range q {
			if root.Left != nil {
				p = append(p, root.Left)
			}
			if root.Right != nil {
				p = append(p, root.Right)
			}
		}
		q = p
	}
	return ans
}
