package e_0003

func numsOfkLevelTreeNode(root *TreeNode, k int) int {
	//return solution1(root, k)
	return solution2(root, k)
}

func solution1(root *TreeNode, k int) int {
	// 递归，深度优先搜索
	if root == nil || k == 0 {
		return 0
	}
	if k == 1 {
		return 1
	}
	return solution1(root.Left, k-1) + solution1(root.Right, k-1)
}

func solution2(root *TreeNode, k int) int {
	// 迭代，广度优先搜索，层序遍历
	if root == nil || k == 0 {
		return 0
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		if k == 1 {
			return len(q)
		}
		var p []*TreeNode
		for _, node := range q {
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		k--
		q = p
	}
	return 0
}
