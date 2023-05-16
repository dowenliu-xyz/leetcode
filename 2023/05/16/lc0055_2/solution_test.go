package lc0055_2

import (
	"fmt"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func Test_canJump(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: `
[2,3,1,1,4]
`,
			want: true,
		},
		{
			input: `
[3,2,1,0,4]
`,
			want: false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			nums, err := inputs.ReadIntSlice(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			if got := canJump(nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
