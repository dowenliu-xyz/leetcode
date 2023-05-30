package inputs

import "github.com/dowenliu-xyz/leetcode/internal/types"

func IntSliceToSingleLinkedList(intSlice []int) *types.SingleLinkedListNode {
	dummy := &types.SingleLinkedListNode{}
	cur := dummy
	for _, v := range intSlice {
		cur.Next = &types.SingleLinkedListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func ReadIntSingleLinkedList(input string) (*types.SingleLinkedListNode, error) {
	ints, err := ReadIntSlice(input)
	if err != nil {
		return nil, err
	}
	return IntSliceToSingleLinkedList(ints), nil
}
