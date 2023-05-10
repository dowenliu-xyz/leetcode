package lc0142_1

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	f, s := head, head
	for {
		if f == nil || f.Next == nil {
			return nil
		}
		f, s = f.Next.Next, s.Next
		if f == s {
			break
		}
	}
	f = head
	for f != s {
		f, s = f.Next, s.Next
	}
	return f
}
