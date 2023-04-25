package lc0104_5

func maxDepth(root *TreeNode) int {
	//return solution1(root)
	return solution2(root)
}

func solution1(root *TreeNode) int {
	// dfs
	if root == nil {
		return 0
	}
	return 1 + max(solution1(root.Left), solution1(root.Right))
}

func solution2(root *TreeNode) int {
	// bfs
	if root == nil {
		return 0
	}
	q, ans := []*TreeNode{root}, 0
	for len(q) > 0 {
		ans++
		var p []*TreeNode
		for _, n := range q {
			if n.Left != nil {
				p = append(p, n.Left)
			}
			if n.Right != nil {
				p = append(p, n.Right)
			}
		}
		q = p
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
