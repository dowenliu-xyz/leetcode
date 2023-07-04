package lc0141

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	s, f := head, head
	for f != nil && f.Next != nil {
		s, f = s.Next, f.Next.Next
		if s == f {
			return true
		}
	}
	return false
}
