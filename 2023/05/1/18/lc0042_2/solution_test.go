package lc0042_2

import (
	"fmt"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func Test_trap(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `[0,1,0,2,1,0,1,3,2,1,2,1]`,
			want:  6,
		},
		{
			input: `[4,2,0,3,2,5]`,
			want:  9,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			height, err := inputs.ReadIntSlice(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			if got := trap(height); got != tt.want {
				t.Errorf("trap(%v) = %v, want %v", height, got, tt.want)
			}
		})
	}
}
