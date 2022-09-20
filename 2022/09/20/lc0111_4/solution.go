package main

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil || root.Right == nil {
		return minDepth(root.Left) + minDepth(root.Right) + 1
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
