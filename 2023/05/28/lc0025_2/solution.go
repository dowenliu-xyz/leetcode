package lc0025_2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	pre, begin := dummy, head
	for {
		end := pre
		for i := 0; i < k; i++ {
			end = end.Next
			if end == nil {
				return dummy.Next
			}
		}
		nextBegin := end.Next
		pre.Next = reverseGroup(begin, nextBegin)
		pre, begin = begin, nextBegin
	}
}

func reverseGroup(head, endNext *ListNode) *ListNode {
	pre := endNext
	for head != endNext {
		pre, head.Next, head = head, pre, head.Next
	}
	return pre
}
