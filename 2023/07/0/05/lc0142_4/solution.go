package lc0142

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			fast = head
			for fast != slow {
				fast, slow = fast.Next, slow.Next
			}
			return fast
		}
	}
	return nil
}

/*
原理解释：
设进行环之前有 a 个节点（即从头节点开始经过 a 步会到达环），环周长为 b 个节点。
第一次相遇时，fast 走了 F 步， slow 走了 S 步。则有
  F = 2S
假设此时 fast 在环内走了 n 圈， slow 走了 m 圈，相遇位置为环入口开始第 c 个位置。则有
  F = a + n*b + c
  S = a + m*b + c
两式相减得：
  F - S = (a + n*b + c) - (a + m*b + c)
  S = (n-m)*b
即 S 刚好是 b 的整数倍。于是可推出 a + c = b ， a + c 刚好是环的周长
此时从相遇位置开始，再走 a 步就可达到环入口位置。
*/
