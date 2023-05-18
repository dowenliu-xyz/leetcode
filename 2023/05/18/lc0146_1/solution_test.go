package lc0146_1

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func TestLRUCache(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `
["LRUCache","put","put","get","put","get","put","get","get","get"]
[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
`,
			want: `[null,null,null,1,null,-1,null,-1,3,4]`,
		},
		{
			input: `
["LRUCache","put","put","get","put","put","get"]
[[2],[2,1],[2,2],[2],[1,1],[4,1],[2]]
`,
			want: `[null,null,null,2,null,null,-1]`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			operations, args, err := readInput(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			var result []any
			c := &LRUCache{}
			for j, operation := range operations {
				op, ok := opMap[operation]
				if !ok {
					t.Fatalf("unexpected operation: %s", operation)
				}
				opResult := op(t, c, args[j])
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

var opMap = map[string]func(t *testing.T, c *LRUCache, args string) any{
	"LRUCache": callConstruct,
	"put":      callPut,
	"get":      callGet,
}

func callConstruct(t *testing.T, c *LRUCache, args string) any {
	capacity, err := strconv.Atoi(strings.TrimSpace(args))
	if err != nil {
		t.Fatalf("invalid args: %v", args)
	}
	*c = Constructor(capacity)
	return nil
}

func callPut(t *testing.T, c *LRUCache, args string) any {
	ints, err := inputs.ReadIntSlice(args)
	if err != nil || len(ints) != 2 {
		t.Fatalf("invalid args: %v", args)
	}
	c.Put(ints[0], ints[1])
	return nil
}

func callGet(t *testing.T, c *LRUCache, args string) any {
	key, err := strconv.Atoi(strings.TrimSpace(args))
	if err != nil {
		t.Fatalf("invalid args: %v", args)
	}
	return c.Get(key)
}
