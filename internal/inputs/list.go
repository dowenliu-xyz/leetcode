package inputs

import "github.com/dowenliu-xyz/leetcode/internal/types"

func IntSliceToSingleLinkedList(intSlice []int) *types.ListNode {
	dummy := &types.ListNode{}
	cur := dummy
	for _, v := range intSlice {
		cur.Next = &types.ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func ReadLinkedList(input string) (*types.ListNode, error) {
	ints, err := ReadIntSlice(input)
	if err != nil {
		return nil, err
	}
	return IntSliceToSingleLinkedList(ints), nil
}
