package of0009

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func TestCQueue(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `
["CQueue","appendTail","deleteHead","deleteHead","deleteHead"]
[[],[3],[],[],[]]
`,
			want: `[null,null,3,-1,-1]`,
		},
		{
			input: `
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
`,
			want: `[null,-1,null,null,5,2]`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			operations, args, err := readInput(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			var result []any
			q := &CQueue{}
			for j, operation := range operations {
				op, ok := opMap[operation]
				if !ok {
					t.Fatalf("unexpected operation: %s", operation)
				}
				opResult := op(t, q, args[j])
				result = append(result, opResult)
			}
			got := sprintResult(result)
			want := strings.ReplaceAll(tt.want, " ", "")
			if got != want {
				t.Errorf("expect %s, got %s", want, got)
			}
		})
	}
}

func readInput(input string) (operations []string, args []string, err error) {
	lines := inputs.ReadNoneBlankLines(input)
	if len(lines) != 2 {
		err = fmt.Errorf("input should have 2 lines, but got %d", len(lines))
		return
	}
	operations, err = inputs.ReadStringSlice(lines[0])
	if err != nil {
		return
	}
	args, err = readArgs(lines[1])
	if err != nil {
		return
	}
	if len(operations) != len(args) {
		err = fmt.Errorf("len(operations): %d, but len(args): %d", len(operations), len(args))
		return
	}
	return
}

func readArgs(input string) ([]string, error) {
	input = strings.ReplaceAll(input, " ", "")
	if !strings.HasPrefix(input, "[[") || !strings.HasSuffix(input, "]]") {
		return nil, fmt.Errorf("input should be wrapped by '[[' and ']]'")
	}
	input = input[2 : len(input)-2]
	ans := strings.Split(input, "],[")
	return ans, nil
}

func sprintResult(result []any) string {
	var sb strings.Builder
	sb.WriteString("[")
	for i := range result {
		if i != 0 {
			sb.WriteString(",")
		}
		if result[i] == nil {
			sb.WriteString("null")
			continue
		}
		sb.WriteString(fmt.Sprintf("%v", result[i]))
	}
	sb.WriteString("]")
	return sb.String()
}

var opMap = map[string]func(t *testing.T, q *CQueue, args string) any{
	"CQueue":     callConstruct,
	"appendTail": callAppendTail,
	"deleteHead": callDeletedHead,
}

func callConstruct(_ *testing.T, q *CQueue, _ string) any {
	*q = Constructor()
	return nil
}

func callAppendTail(t *testing.T, q *CQueue, args string) any {
	value, err := strconv.Atoi(args)
	if err != nil {
		t.Fatalf("invalid args: %v", args)
	}
	q.AppendTail(value)
	return nil
}

func callDeletedHead(_ *testing.T, q *CQueue, _ string) any {
	return q.DeleteHead()
}
