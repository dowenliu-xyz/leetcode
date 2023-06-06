package lc0141_1

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func TestHasCycle(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: `
[3,2,0,-4]
1
`,
			want: true,
		},
		{
			input: `
[1,2]
0
`,
			want: true,
		},
		{
			input: `
[1]
-1
`,
			want: false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			head, err := readIntList(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			got := hasCycle(head)
			if got != tt.want {
				t.Errorf("expect %t, got %t", tt.want, got)
			}
		})
	}
}

func readIntList(input string) (*ListNode, error) {
	lines := inputs.ReadNoneBlankLines(input)
	if len(lines) != 2 {
		return nil, fmt.Errorf("expect 2 lines, got %d", len(lines))
	}
	vals := strings.ReplaceAll(lines[0], " ", "")
	if vals == "[]" {
		return nil, nil
	}
	intSlice, err := inputs.ReadIntSlice(vals)
	if err != nil {
		return nil, err
	}
	pos, err := strconv.Atoi(lines[1])
	if err != nil {
		return nil, err
	}
	return intSliceToList(intSlice, pos), nil
}

func intSliceToList(intSlice []int, pos int) *ListNode {
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
	return dummy.Next
}
