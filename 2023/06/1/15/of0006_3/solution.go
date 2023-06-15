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
		stack, head = append(stack, head.Val), head.Next
	}
	for l, r := 0, len(stack)-1; l < r; l, r = l+1, r-1 {
		stack[l], stack[r] = stack[r], stack[l]
	}
	return stack
}
