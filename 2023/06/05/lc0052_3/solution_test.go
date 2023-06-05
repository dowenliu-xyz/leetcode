package lc0052

import (
	"fmt"
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
