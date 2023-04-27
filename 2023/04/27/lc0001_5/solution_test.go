package lc0001_5

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `
[2,7,11,15]
9
`,
			want: `[0,1]`,
		},
		{
			input: `
[3,2,4]
6
`,
			want: `[1,2]`,
		},
		{
			input: `
[3,3]
6
`,
			want: `[0,1]`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			lines := inputs.ReadNoneBlankLines(tt.input)
			if len(lines) != 2 {
				t.Fatalf("expect 2 lines, got %d", len(lines))
			}
			nums, err := inputs.ReadIntSlice(lines[0])
			if err != nil {
				t.Fatal(err)
			}
			target, err := strconv.ParseInt(strings.TrimSpace(lines[1]), 10, 64)
			if err != nil {
				t.Fatal(err)
			}
			got := twoSum(nums, int(target))
			sort.Ints(got)
			result := fmt.Sprintf("%v", got)
			result = strings.ReplaceAll(result, " ", ",")
			if tt.want != result {
				t.Fatalf("expect %s, got %s", tt.want, result)
			}
		})
	}
}
