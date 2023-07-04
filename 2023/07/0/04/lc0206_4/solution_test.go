package lc0206

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `[1,2,3,4,5]`,
			want:  `[5,4,3,2,1]`,
		},
		{
			input: `[1,2]`,
			want:  `[2,1]`,
		},
		{
			input: `[]`,
			want:  `[]`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			head, err := readIntList(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			got := reverseList(head)
			result := sprintList(got)
			if result != tt.want {
				t.Errorf("want: %v, got: %v", tt.want, result)
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
