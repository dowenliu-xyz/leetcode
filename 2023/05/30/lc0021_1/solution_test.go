package lc0021_1

import (
	"fmt"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
	"github.com/dowenliu-xyz/leetcode/internal/outputs"
	"github.com/dowenliu-xyz/leetcode/internal/types"
)

func Test_mergeTwoLists(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `
[1,2,4]
[1,3,4]
`,
			want: `[1,1,2,3,4,4]`,
		},
		{
			input: `
[]
[]
`,
			want: `[]`,
		},
		{
			input: `
[]
[0]
`,
			want: `[0]`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			list1, list2, err := readInput(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			got := mergeTwoLists(list1, list2)
			if got := outputs.PrintSingleLinkedList(toInternal(got)); got != tt.want {
				t.Errorf("expect %s, got %s", tt.want, got)
			}
		})
	}
}

func readInput(input string) (*ListNode, *ListNode, error) {
	lines := inputs.ReadNoneBlankLines(input)
	if len(lines) != 2 {
		return nil, nil, fmt.Errorf("input should have 2 lines")
	}
	list1, err := inputs.ReadIntSingleLinkedList(lines[0])
	if err != nil {
		return nil, nil, err
	}
	list2, err := inputs.ReadIntSingleLinkedList(lines[1])
	if err != nil {
		return nil, nil, err
	}
	return fromInternal(list1), fromInternal(list2), nil
}

func fromInternal(node *types.SingleLinkedListNode) *ListNode {
	if node == nil {
		return nil
	}
	return &ListNode{
		Val:  node.Val,
		Next: fromInternal(node.Next),
	}
}

func toInternal(node *ListNode) *types.SingleLinkedListNode {
	if node == nil {
		return nil
	}
	return &types.SingleLinkedListNode{
		Val:  node.Val,
		Next: toInternal(node.Next),
	}
}
