package lc0023

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
	"github.com/dowenliu-xyz/leetcode/internal/outputs"
	"github.com/dowenliu-xyz/leetcode/internal/types"
)

func Test_mergeKLists(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `[[1,4,5],[1,3,4],[2,6]]`,
			want:  `[1,1,2,3,4,4,5,6]`,
		},
		{
			input: `[]`,
			want:  `[]`,
		},
		{
			input: `[[1]]`,
			want:  `[1]`,
		},
		{
			input: `[[]]`,
			want:  `[]`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			done := make(chan struct{})
			go func() {
				defer func() {
					done <- struct{}{}
				}()
				lists, err := readInput(tt.input)
				if err != nil {
					t.Error(err)
					return
				}
				got := mergeKLists(lists)
				want := strings.ReplaceAll(tt.want, " ", "")
				if got := outputs.SprintLinkedList(toInternal(got)); got != want {
					t.Errorf("mergeKLists() = %v, want %v", got, want)
				}
			}()
			select {
			case <-done:
			case <-time.After(time.Second):
				t.Error("timeout")
			}
		})
	}
}

func readInput(input string) ([]*ListNode, error) {
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ReplaceAll(input, "\n", "")
	if input == "[]" {
		return nil, nil
	}
	if !strings.HasPrefix(input, "[[") || !strings.HasSuffix(input, "]]") {
		return nil, fmt.Errorf("invalid input")
	}
	input = strings.TrimPrefix(input, "[[")
	input = strings.TrimSuffix(input, "]]")
	listStrs := strings.Split(input, "],[")
	lists := make([]*ListNode, len(listStrs))
	for _, listStr := range listStrs {
		internalList, err := inputs.ReadLinkedList(fmt.Sprintf("[%s]", listStr))
		if err != nil {
			return nil, err
		}
		lists = append(lists, fromInternal(internalList))
	}
	return lists, nil
}

func fromInternal(internal *types.ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for internal != nil {
		cur.Next = &ListNode{Val: internal.Val}
		cur = cur.Next
		internal = internal.Next
	}
	return dummy.Next
}

func toInternal(node *ListNode) *types.ListNode {
	dummy := &types.ListNode{}
	cur := dummy
	for node != nil {
		cur.Next = &types.ListNode{Val: node.Val}
		cur = cur.Next
		node = node.Next
	}
	return dummy.Next
}
