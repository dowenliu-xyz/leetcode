package lc0052_1

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{
			input: 4,
			want:  2,
		},
		{
			input: 1,
			want:  1,
		},
		{
			input: 7,
			want:  40,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if got := totalNQueens(tt.input); got != tt.want {
				t.Fatalf("expect %d, got %d", tt.want, got)
			}
		})
	}
}

func sprintResult(result [][]string) string {
	square := make([]string, 0, len(result))
	for _, s := range result {
		square = append(square, fmt.Sprintf(`["%s"]`, strings.Join(s, `","`)))
	}
	sort.Slice(square, func(i, j int) bool {
		return strings.Compare(square[i], square[j]) >= 0
	})
	return fmt.Sprintf("[%s]", strings.Join(square, ","))
}
