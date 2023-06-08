package of0006

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	var stack []int
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}
	l, r := 0, len(stack)-1
	for l < r {
		stack[l], stack[r] = stack[r], stack[l]
		l++
		r--
	}
	return stack
}
