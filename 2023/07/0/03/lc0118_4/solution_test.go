package lc0118

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{
			input: 5,
			want:  `[[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]`,
		},
		{
			input: 1,
			want:  `[[1]]`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			result := generate(tt.input)
			got := printResult(result)
			if tt.want != got {
				t.Fatalf("expect %s, got %s", tt.want, got)
			}
		})
	}
}

func printResult(result [][]int) string {
	lines := make([]string, 0, len(result))
	for _, ints := range result {
		nums := make([]string, 0, len(ints))
		for _, i := range ints {
			nums = append(nums, strconv.FormatInt(int64(i), 10))
		}
		lines = append(lines, fmt.Sprintf("[%s]", strings.Join(nums, ",")))
	}
	return fmt.Sprintf("[%s]", strings.Join(lines, ","))
}
