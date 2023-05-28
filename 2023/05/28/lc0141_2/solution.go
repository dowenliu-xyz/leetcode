package lc0141_2

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
		s = s.Next
		f = f.Next.Next
		if f == s {
			return true
		}
	}
	return false
}
