package lc0455

import (
	"fmt"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func TestFindContentChildren(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `
[1,2,3]
[1,1]
`,
			want: 1,
		},
		{
			input: `
[1,2]
[1,2,3]
`,
			want: 2,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			g, s := parseInput(t, tt.input)
			got := findContentChildren(g, s)
			if got != tt.want {
				t.Errorf("expect %d, got %d", tt.want, got)
			}
		})
	}
}

func parseInput(t *testing.T, input string) (g, s []int) {
	t.Helper()

	lines := inputs.ReadNoneBlankLines(input)
	if len(lines) != 2 {
		t.Fatalf("expect 2 lines, got %d", len(lines))
	}
	var err error
	g, err = inputs.ReadIntSlice(lines[0])
	if err != nil {
		t.Fatal(err)
	}
	s, err = inputs.ReadIntSlice(lines[1])
	if err != nil {
		t.Fatal(err)
	}
	return
}
