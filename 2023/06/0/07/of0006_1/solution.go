package of0006

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	var pre *ListNode
	n := 0
	for head != nil {
		n++
		head.Next, head, pre = pre, head.Next, head
	}
	head = pre
	ans := make([]int, 0, n)
	for head != nil {
		ans = append(ans, head.Val)
		head = head.Next
	}
	return ans
}
