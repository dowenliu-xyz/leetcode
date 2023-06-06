package lc0024_1

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre, cur := dummy, head
	for cur != nil && cur.Next != nil {
		next := cur.Next
		pre.Next, cur.Next, next.Next = next, next.Next, cur
		pre, cur = cur, cur.Next
	}
	return dummy.Next
}
