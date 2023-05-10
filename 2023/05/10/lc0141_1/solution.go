package lc0141_1

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	f, s := head, head
	for f != nil && f.Next != nil {
		f, s = f.Next.Next, s.Next
		if f == s {
			return true
		}
	}
	return false
}
