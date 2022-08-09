package lc0104_1

func maxDepth(root *TreeNode) int {
	//return solution1(root)
	return solution2(root)
}

func solution1(root *TreeNode) int {
	// 递归，深度优先搜索
	if root == nil {
		return 0
	}
	return 1 + maxInt(solution1(root.Left), solution1(root.Right))
}

func solution2(root *TreeNode) int {
	// 迭代，广度优先搜索
	if root == nil {
		return 0
	}
	res := 0
	q := []*TreeNode{root}
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

func maxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}
