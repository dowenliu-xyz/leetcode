package main

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
	for l1 != nil || l2 != nil {
		node.Next = &ListNode{Val: carry}
		node = node.Next
		if l1 != nil {
			node.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			node.Val += l2.Val
			l2 = l2.Next
		}
		carry, node.Val = node.Val/10, node.Val%10
	}
	if carry > 0 {
		node.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}
