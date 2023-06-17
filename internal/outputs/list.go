package outputs

import (
	"strconv"
	"strings"

	"github.com/dowenliu-xyz/leetcode/internal/types"
)

func SprintLinkedList(head *types.ListNode) string {
	bu := strings.Builder{}
	bu.WriteByte('[')
	for head != nil {
		bu.WriteString(strconv.Itoa(head.Val))
		head = head.Next
		if head != nil {
			bu.WriteByte(',')
		}
	}
	bu.WriteByte(']')
	return bu.String()
}
