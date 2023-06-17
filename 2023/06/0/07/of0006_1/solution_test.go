package of0006

import (
	"fmt"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
	"github.com/dowenliu-xyz/leetcode/internal/outputs"
	"github.com/dowenliu-xyz/leetcode/internal/types"
)

func Test_reversePrint(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"[1,3,2]", "[2,3,1]"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			head, err := readInput(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			got := reversePrint(head)
			want := strings.ReplaceAll(tt.want, " ", "")
			if got := outputs.SprintIntSlice(got); got != want {
				t.Errorf("reversePrint() = %v, want %v", got, want)
			}
		})
	}
}

func readInput(input string) (*ListNode, error) {
	head, err := inputs.ReadLinkedList(input)
	if err != nil {
		return nil, err
	}
	return fromInternal(head), nil
}

func fromInternal(head *types.ListNode) *ListNode {
	dummy := &ListNode{}
	pre := dummy
	for head != nil {
		pre.Next = &ListNode{Val: head.Val}
		pre, head = pre.Next, head.Next
	}
	return dummy.Next
}
