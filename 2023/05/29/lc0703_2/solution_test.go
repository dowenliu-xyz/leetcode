package lc0703_2

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func Test_KthLargest(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `
["KthLargest", "add", "add", "add", "add", "add"]
[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
`,
			want: `[null, 4, 5, 5, 8, 8]`,
		},
		{
			input: `
["KthLargest","add","add","add","add","add"]
[[1,[]],[-3],[-2],[-4],[0],[4]]
`,
			want: `[null, -3, -2, -2, 0, 4]`,
		},
		{
			input: `
["KthLargest","add","add","add","add","add","add","add","add","add","add","add","add","add","add","add","add","add","add","add","add","add","add","add","add","add","add","add"]
[[7,[-10,1,3,1,4,10,3,9,4,5,1]],[3],[2],[3],[1],[2],[4],[5],[5],[6],[7],[7],[8],[2],[3],[1],[1],[1],[10],[11],[5],[6],[2],[4],[7],[8],[5],[6]]
`,
			want: `[null,3,3,3,3,3,3,4,4,4,5,5,5,5,5,5,5,5,6,7,7,7,7,7,7,7,7,7]`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			operations, args, err := readInput(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			var result []any
			p := &KthLargest{}
			for j, operation := range operations {
				op, ok := oppMap[operation]
				if !ok {
					t.Fatalf("unexpected operation: %s", operation)
				}
				opResult := op(t, p, args[j])
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

var oppMap = map[string]func(t *testing.T, p *KthLargest, args string) any{
	"KthLargest": callConstruct,
	"add":        callAdd,
}

func callConstruct(t *testing.T, p *KthLargest, args string) any {
	// 3,[4,5,8,2]
	args = strings.ReplaceAll(args, " ", "")
	parts := strings.SplitN(args, ",", 2)
	if len(parts) != 2 {
		t.Fatalf("invalid args: %s", args)
	}
	k, err := strconv.Atoi(parts[0])
	if err != nil {
		t.Fatalf("invalid args: %s", args)
	}
	nums, err := inputs.ReadIntSlice(parts[1])
	if err != nil {
		t.Fatalf("invalid args: %s", args)
	}
	*p = Constructor(k, nums)
	return nil
}

func callAdd(t *testing.T, p *KthLargest, args string) any {
	// 3
	args = strings.ReplaceAll(args, " ", "")
	val, err := strconv.Atoi(args)
	if err != nil {
		t.Fatalf("invalid args: %s", args)
	}
	return p.Add(val)
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
	// [[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
	input = strings.ReplaceAll(input, " ", "")
	// [[3,[4,5,8,2]],[3],[5],[10],[9],[4]]
	if !strings.HasPrefix(input, "[") || !strings.HasSuffix(input, "]") {
		return nil, fmt.Errorf("input should be wrapped by '[' and ']'")
	}
	input = input[1 : len(input)-1] // remove '[' and ']'
	// [3,[4,5,8,2]],[3],[5],[10],[9],[4]
	input = input[1 : len(input)-1] // remove '[' and ']'
	// 3,[4,5,8,2]],[3],[5],[10],[9],[4
	ans := strings.Split(input, "],[")
	// 3,[4,5,8,2] | 3 | 5 | 10 | 9 | 4
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
