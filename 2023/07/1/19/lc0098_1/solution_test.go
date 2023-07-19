package lc0098

import (
	"fmt"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func Test_isValidBST(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"[2,1,3]", true},
		{"[5,1,4,null,null,3,6]", false},
		{"[2,2,2]", false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			root, err := inputs.LevelOrderRestore(tt.input)
			if err != nil {
				t.Fatalf("failed to restore tree from %s: %v", tt.input, err)
			}
			if got := isValidBST(root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
