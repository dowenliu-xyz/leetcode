package lc0021

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	pre := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			pre.Next, list1 = list1, list1.Next
		} else {
			pre.Next, list2 = list2, list2.Next
		}
		pre = pre.Next
	}
	if list1 != nil {
		pre.Next = list1
	}
	if list2 != nil {
		pre.Next = list2
	}
	return dummy.Next
}
