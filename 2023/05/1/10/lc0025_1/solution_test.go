package lc0025_1

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func Test_reverseKGroup(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `
[1,2,3,4,5]
2
`,
			want: `[2,1,4,3,5]`,
		},
		{
			input: `
[1,2,3,4,5]
3
`,
			want: `[3,2,1,4,5]`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			head, k, err := readInput(tt.input)
			if err != nil {
				t.Fatalf("readInput(%s) failed: %v", tt.input, err)
			}
			got := reverseKGroup(head, k)
			if got := sprintList(got); got != tt.want {
				t.Errorf("reverseKGroup(%s, %d) = %s, want %s", tt.input, k, got, tt.want)
			}
		})
	}
}

func readInput(input string) (*ListNode, int, error) {
	lines := inputs.ReadNoneBlankLines(input)
	if len(lines) != 2 {
		return nil, 0, fmt.Errorf("input should have 2 lines")
	}
	head, err := readIntList(lines[0])
	if err != nil {
		return nil, 0, err
	}
	k, err := strconv.Atoi(lines[1])
	if err != nil {
		return nil, 0, err
	}
	return head, k, nil
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
