package lc0101

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	//return solution1(root)
	return solution2(root)
}

func solution1(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

func solution2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	q := [][]*TreeNode{{root.Left, root.Right}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		l, r := p[0], p[1]
		if l == nil && r == nil {
			continue
		}
		if l == nil || r == nil {
			return false
		}
		if l.Val != r.Val {
			return false
		}
		q = append(q, []*TreeNode{l.Left, r.Right}, []*TreeNode{l.Right, r.Left})
	}
	return true
}
