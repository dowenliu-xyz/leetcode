package lc0045_3

import (
	"fmt"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func Test_jump(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `
[2,3,1,1,4]
`,
			want: 2,
		},
		{
			input: `
[2,3,0,1,4]
`,
			want: 2,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			nums, err := inputs.ReadIntSlice(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			if got := jump(nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
