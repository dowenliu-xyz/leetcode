package lc0142

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	f, s := head, head
	for f != nil && f.Next != nil {
		f, s = f.Next.Next, s.Next
		if f == s {
			f = head
			for f != s {
				f, s = f.Next, s.Next
			}
			return s
		}
	}
	return nil
}
