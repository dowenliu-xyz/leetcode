package lc0092

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	pre, start := dummy, dummy
	for i := 0; i < left; i++ {
		pre, start = start, start.Next
	}
	end := start
	for i := left; i < right; i++ {
		end = end.Next
	}
	pre.Next = reverse(start, end.Next)
	return dummy.Next
}

func reverse(head, next *ListNode) *ListNode {
	pre := next
	for head != next {
		pre, head, head.Next = head, head.Next, pre
	}
	return pre
}
