package lc0002_1

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	carry, node := 0, dummy
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + carry
		carry, val = val/10, val%10
		node.Next = &ListNode{Val: val}
		node = node.Next
		l1, l2 = l1.Next, l2.Next
	}
	for l1 != nil {
		val := l1.Val + carry
		carry, val = val/10, val%10
		node.Next = &ListNode{Val: val}
		node = node.Next
		l1 = l1.Next
	}
	for l2 != nil {
		val := l2.Val + carry
		carry, val = val/10, val%10
		node.Next = &ListNode{Val: val}
		node = node.Next
		l2 = l2.Next
	}
	if carry > 0 {
		node.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}
