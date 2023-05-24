package lc0056_3

import (
	"fmt"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
	"github.com/dowenliu-xyz/leetcode/internal/outputs"
)

func Test_merge(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `[[1,3],[2,6],[8,10],[15,18]]`,
			want:  `[[1,6],[8,10],[15,18]]`,
		},
		{
			input: `[[1,4],[4,5]]`,
			want:  `[[1,5]]`,
		},
		{
			input: `[[1,4],[2,3]]`,
			want:  `[[1,4]]`,
		},
		{
			input: `[[1,4],[0,4]]`,
			want:  `[[0,4]]`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			intervals, err := inputs.ReadIntSquareInLine(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			got := merge(intervals)
			want := strings.ReplaceAll(tt.want, " ", "")
			if got := outputs.SprintIntMatrix(got); got != want {
				t.Errorf("merge(%v) = %v, want %v", intervals, got, tt.want)
			}
		})
	}
}
