package e_0002

func numsOfNoChildNode(root *TreeNode) int {
	//return solution1(root)
	return solution2(root)
}

func solution1(root *TreeNode) int {
	// 递归 深度优先搜索
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return solution1(root.Left) + solution1(root.Right)
}

func solution2(root *TreeNode) int {
	// 迭代 广度优先搜索 层序遍历
	if root == nil {
		return 0
	}
	res, q := 0, []*TreeNode{root}
	for len(q) > 0 {
		var p []*TreeNode
		for _, node := range q {
			if node.Left == nil && node.Right == nil {
				res++
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
	return res
}
