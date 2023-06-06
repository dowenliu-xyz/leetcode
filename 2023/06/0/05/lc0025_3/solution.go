package lc0025

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	var reverseGroup func(head, next *ListNode) *ListNode
	reverseGroup = func(head, next *ListNode) *ListNode {
		pre := next
		for head != next {
			pre, head.Next, head = head, pre, head.Next
		}
		return pre
	}
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
