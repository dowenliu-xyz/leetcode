package main

func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh, rh := height(root.Left), height(root.Right)
	if lh == -1 || rh == -1 || abs(lh-rh) > 1 {
		return -1
	}
	return max(lh, rh) + 1
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
