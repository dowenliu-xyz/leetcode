package lc0225

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func TestMyStack(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `
["MyStack", "push", "push", "top", "pop", "empty"]
[[], [1], [2], [], [], []]
`,
			want: `[null, null, null, 2, 2, false]`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			operations, args, err := readInput(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			var result []any
			q := &MyStack{}
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

var opMap = map[string]func(t *testing.T, q *MyStack, args string) any{
	"MyStack": callConstruct,
	"push":    callPush,
	"top":     callTop,
	"pop":     callPop,
	"empty":   callEmpty,
}

func callConstruct(_ *testing.T, q *MyStack, _ string) any {
	*q = Constructor()
	return nil
}

func callPush(t *testing.T, q *MyStack, args string) any {
	value, err := strconv.Atoi(args)
	if err != nil {
		t.Fatalf("invalid args: %v", args)
	}
	q.Push(value)
	return nil
}

func callTop(_ *testing.T, q *MyStack, _ string) any {
	return q.Top()
}

func callPop(_ *testing.T, q *MyStack, _ string) any {
	return q.Pop()
}

func callEmpty(_ *testing.T, q *MyStack, _ string) any {
	return q.Empty()
}
