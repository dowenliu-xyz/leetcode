package lc0322_2

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func TestCoinChange(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `
[1,2,5]
11
`,
			want: 3,
		},
		{
			input: `
[2]
3
`,
			want: -1,
		},
		{
			input: `
[1]
0
`,
			want: 0,
		},
		{
			input: `
[186,419,83,408]
6249
`,
			want: 20,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			coins, amount := parseInput(t, tt.input)
			got := coinChange(coins, amount)
			if got != tt.want {
				t.Errorf("expect %d, got %d", tt.want, got)
			}
		})
	}
}

func parseInput(t *testing.T, input string) (coins []int, amount int) {
	t.Helper()

	lines := inputs.ReadNoneBlankLines(input)
	if len(lines) != 2 {
		t.Fatalf("expect 2 lines, got %d", len(lines))
	}
	var err error
	coins, err = inputs.ReadIntSlice(lines[0])
	if err != nil {
		t.Fatal(err)
	}
	i, err := strconv.ParseInt(lines[1], 10, 64)
	if err != nil {
		t.Fatal(err)
	}
	amount = int(i)
	return
}
