package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	carry, cur := 0, dummy
	for l1 != nil || l2 != nil {
		val := carry
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		val, carry = val%10, val/10
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}
	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}
