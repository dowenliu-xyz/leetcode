package e_0001

func numOfTreeNode(root *TreeNode) int {
	//return solution1(root)
	return solution2(root)
}

func solution1(root *TreeNode) int {
	// 递归 深度优先搜索
	if root == nil {
		return 0
	}
	return solution1(root.Left) + solution1(root.Right) + 1
}

func solution2(root *TreeNode) int {
	// 迭代。广度优先搜索，层序遍历
	if root == nil {
		return 0
	}
	res, q := 0, []*TreeNode{root}
	for len(q) > 0 {
		var p []*TreeNode
		for _, node := range q {
			res++
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
