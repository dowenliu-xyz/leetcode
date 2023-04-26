package of0051_1

import (
	"fmt"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func TestReversePairs(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `[7,5,6,4]`,
			want:  5,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			nums, err := inputs.ReadIntSlice(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			got := reversePairs(nums)
			if tt.want != got {
				t.Fatalf("expect %d, got %d", tt.want, got)
			}
		})
	}
}
