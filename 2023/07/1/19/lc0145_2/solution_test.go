package lc0145

import (
	"fmt"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
	"github.com/dowenliu-xyz/leetcode/internal/outputs"
)

func Test_postorderTraversal(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"[1,null,2,3]", "[3,2,1]"},
		{"[]", "[]"},
		{"[1]", "[1]"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			root, err := inputs.LevelOrderRestore(tt.input)
			if err != nil {
				t.Fatalf("unexpected input: %s: %v", tt.input, err)
			}
			got := postorderTraversal(root)
			if !match(outputs.SprintIntSlice(got), tt.want) {
				t.Fatalf("postorderTraversal(%v) = %v; want %s", tt.input, got, tt.want)
			}
		})
	}
}

func match(got, want string) bool {
	got = strings.ReplaceAll(got, " ", "")
	want = strings.ReplaceAll(want, " ", "")
	return got == want
}
