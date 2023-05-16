package lc0011_2

import (
	"fmt"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func Test_maxArea(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `
[1,8,6,2,5,4,8,3,7]
`,
			want: 49,
		},
		{
			input: `
[1,1]
`,
			want: 1,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			height, err := inputs.ReadIntSlice(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			if got := maxArea(height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
