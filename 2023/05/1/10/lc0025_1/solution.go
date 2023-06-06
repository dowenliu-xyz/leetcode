package lc0025_1

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
			if end.Next == nil {
				return dummy.Next
			}
			end = end.Next
		}
		next := end.Next
		pre.Next = reverse(begin, next)
		pre, begin = begin, next
	}
}

func reverse(head, endNext *ListNode) *ListNode {
	pre, cur := endNext, head
	for cur != endNext {
		pre, cur.Next, cur = cur, pre, cur.Next
	}
	return pre
}
