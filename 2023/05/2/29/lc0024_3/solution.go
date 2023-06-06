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
	for head != nil && head.Next != nil {
		pre.Next, head.Next, head.Next.Next = head.Next, head.Next.Next, head
		pre, head = head, head.Next
	}
	return dummy.Next
}
