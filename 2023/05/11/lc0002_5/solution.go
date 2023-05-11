package lc0002_5

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	v := 0
	pre := dummy
	for {
		if l1 == nil && l2 == nil {
			break
		}
		if l1 != nil {
			v += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v += l2.Val
			l2 = l2.Next
		}
		pre.Next = &ListNode{Val: v % 10}
		pre, v = pre.Next, v/10
	}
	if v != 0 {
		pre.Next = &ListNode{Val: v}
	}
	return dummy.Next
}
