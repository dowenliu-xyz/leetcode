package lc0206_3

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	pre, cur := (*ListNode)(nil), head
	for cur != nil {
		pre, cur.Next, cur = cur, pre, cur.Next
	}
	return pre
}
