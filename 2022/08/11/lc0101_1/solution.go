package lc0101_1

func isSymmetric(root *TreeNode) bool {
	//return solution1(root)
	return solution2(root)
}

func solution1(root *TreeNode) bool {
	// 递归
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return isMirror(root1.Left, root2.Right) && isMirror(root1.Right, root2.Left)
}

func solution2(root *TreeNode) bool {
	// 迭代
	if root == nil {
		return true
	}
	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) > 1 {
		node1, node2 := queue[0], queue[1]
		queue = queue[2:]
		if (node1 == nil && node2 != nil) || (node1 != nil && node2 == nil) {
			return false
		}
		if node1 == nil && node2 == nil {
			continue
		}
		if node1.Val != node2.Val {
			return false
		}
		queue = append(queue, node1.Left, node2.Right, node1.Right, node2.Left)
	}
	return len(queue) == 0
}
