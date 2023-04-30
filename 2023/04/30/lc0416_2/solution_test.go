package lc0416_2

import (
	"fmt"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func TestCanPartition(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: `[1,5,11,5]`,
			want:  true,
		},
		{
			input: `[1,2,3,5]`,
			want:  false,
		},
		{
			input: `[2,2,3,5]`,
			want:  false,
		},
		{
			input: `[3,3,3,4,5]`,
			want:  true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			nums, err := inputs.ReadIntSlice(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			got := canPartition(nums)
			if tt.want != got {
				t.Fatalf("expect %t, got %t", tt.want, got)
			}
		})
	}
}
