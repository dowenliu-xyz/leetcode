package lc0104_2

func maxDepth(root *TreeNode) int {
	//return solution1(root)
	return solution2(root)
}

func solution1(root *TreeNode) int {
	// 递归，深度优先搜索
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return max(solution1(root.Left), solution1(root.Right)) + 1
}

func solution2(root *TreeNode) int {
	// 迭代，广度优先搜索
	if root == nil {
		return 0
	}
	res, q := 0, []*TreeNode{root}
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
		res++
		q = p
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
