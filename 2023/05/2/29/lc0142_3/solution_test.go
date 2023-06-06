package lc0142_3

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func TestDetectCycle(t *testing.T) {
	tests := []struct {
		input string
	}{
		{
			input: `
[3,2,0,-4]
1
`,
		},
		{
			input: `
[1,2]
0
`,
		},
		{
			input: `
[1]
-1
`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			head, want, err := readIntList(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			got := detectCycle(head)
			if got != want {
				t.Errorf("expect %#v, got %#v", want, got)
			}
		})
	}
}

func readIntList(input string) (*ListNode, *ListNode, error) {
	lines := inputs.ReadNoneBlankLines(input)
	if len(lines) != 2 {
		return nil, nil, fmt.Errorf("expect 2 lines, got %d", len(lines))
	}
	vals := strings.ReplaceAll(lines[0], " ", "")
	if vals == "[]" {
		return nil, nil, nil
	}
	intSlice, err := inputs.ReadIntSlice(vals)
	if err != nil {
		return nil, nil, err
	}
	pos, err := strconv.Atoi(lines[1])
	if err != nil {
		return nil, nil, err
	}
	head, want := intSliceToList(intSlice, pos)
	return head, want, nil
}

func intSliceToList(intSlice []int, pos int) (*ListNode, *ListNode) {
	var atPos *ListNode
	dummy := &ListNode{}
	pre, cur := (*ListNode)(nil), dummy
	for i, v := range intSlice {
		cur.Next = &ListNode{Val: v}
		pre, cur = cur, cur.Next
		if i == pos {
			atPos = cur
		}
	}
	if pre != nil {
		pre.Next = atPos
	}
	return dummy.Next, atPos
}
