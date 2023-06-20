package lc0092

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
	"github.com/dowenliu-xyz/leetcode/internal/outputs"
	"github.com/dowenliu-xyz/leetcode/internal/types"
)

func Test_reverseBetween(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `
[1,2,3,4,5]
2
4
`,
			want: `[1,4,3,2,5]`,
		},
		{
			input: `
[5]
1
1
`,
			want: `[5]`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			head, left, right, err := readInput(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			got := reverseBetween(head, left, right)
			want := strings.ReplaceAll(tt.want, " ", "")
			if got := outputs.SprintLinkedList(toInternal(got)); got != want {
				t.Fatalf("want: %s, got: %s", want, got)
			}
		})
	}
}

func readInput(input string) (*ListNode, int, int, error) {
	lines := inputs.ReadNoneBlankLines(input)
	if len(lines) != 3 {
		return nil, 0, 0, fmt.Errorf("input should have 3 lines")
	}
	head, err := inputs.ReadLinkedList(lines[0])
	if err != nil {
		return nil, 0, 0, err
	}
	left, err := strconv.Atoi(strings.TrimSpace(lines[1]))
	if err != nil {
		return nil, 0, 0, err
	}
	right, err := strconv.Atoi(strings.TrimSpace(lines[2]))
	if err != nil {
		return nil, 0, 0, err
	}
	return fromInternal(head), left, right, nil
}

func fromInternal(head *types.ListNode) *ListNode {
	if head == nil {
		return nil
	}
	return &ListNode{
		Val:  head.Val,
		Next: fromInternal(head.Next),
	}
}

func toInternal(head *ListNode) *types.ListNode {
	if head == nil {
		return nil
	}
	return &types.ListNode{
		Val:  head.Val,
		Next: toInternal(head.Next),
	}
}
