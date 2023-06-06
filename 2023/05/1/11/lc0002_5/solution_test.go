package lc0002_5

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func Test_addTwoNumbers(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `
[2,4,3]
[5,6,4]
`,
			want: `[7,0,8]`,
		},
		{
			input: `
[0]
[0]
`,
			want: `[0]`,
		},
		{
			input: `
[9,9,9,9,9,9,9]
[9,9,9,9]
`,
			want: `[8,9,9,9,0,0,0,1]`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			l1, l2, err := readInput(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			got := addTwoNumbers(l1, l2)
			if got := sprintList(got); got != tt.want {
				t.Fatalf("expect %s, got %s", tt.want, got)
			}
		})
	}
}

func readInput(input string) (*ListNode, *ListNode, error) {
	lines := inputs.ReadNoneBlankLines(input)
	if len(lines) != 2 {
		return nil, nil, fmt.Errorf("input should have 2 lines")
	}
	l1, err := readIntList(lines[0])
	if err != nil {
		return nil, nil, err
	}
	l2, err := readIntList(lines[1])
	if err != nil {
		return nil, nil, err
	}
	return l1, l2, nil
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
