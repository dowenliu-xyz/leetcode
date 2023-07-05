package lc0024

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for {
		if head == nil {
			return dummy.Next
		}
		second := head.Next
		if second == nil {
			return dummy.Next
		}
		pre, pre.Next, head, head.Next, second.Next = head, second, second.Next, second.Next, head
	}
}
