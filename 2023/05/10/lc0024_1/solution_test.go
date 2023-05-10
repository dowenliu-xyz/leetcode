package lc0024_1

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func Test_swapPairs(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `[1,2,3,4]`,
			want:  `[2,1,4,3]`,
		},
		{
			input: `[]`,
			want:  `[]`,
		},
		{
			input: `[1]`,
			want:  `[1]`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			head, err := readIntList(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			got := swapPairs(head)
			if gotStr := sprintList(got); gotStr != tt.want {
				t.Errorf("swapPairs() = %v, want %v", gotStr, tt.want)
			}
		})
	}
}

func readIntList(input string) (*ListNode, error) {
	input = strings.ReplaceAll(input, " ", "")
	if input == "[]" {
		return nil, nil
	}
	intSlice, err := inputs.ReadIntSlice(input)
	if err != nil {
		return nil, err
	}
	return intSliceToList(intSlice), nil
}

func intSliceToList(intSlice []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, i := range intSlice {
		current.Next = &ListNode{Val: i}
		current = current.Next
	}
	return dummy.Next
}

func sprintList(head *ListNode) string {
	if head == nil {
		return "[]"
	}
	result := bytes.Buffer{}
	result.WriteString("[")
	for head != nil {
		result.WriteString(fmt.Sprintf("%d", head.Val))
		head = head.Next
		if head != nil {
			result.WriteString(",")
		}
	}
	result.WriteString("]")
	return result.String()
}
