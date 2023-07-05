package lc0025

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for {
		end := pre
		for i := 0; i < k; i++ {
			end = end.Next
			if end == nil {
				return dummy.Next
			}
		}
		next := end.Next
		pre.Next = reverseGroup(head, next)
		pre, head = head, next
	}
}

func reverseGroup(head, next *ListNode) *ListNode {
	pre := next
	for head != next {
		pre, head, head.Next = head, head.Next, pre
	}
	return pre
}
