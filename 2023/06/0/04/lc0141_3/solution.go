package lc0141

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
		f = f.Next.Next
		s = s.Next
		if f == s {
			return true
		}
	}
	return false
}
