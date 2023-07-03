package lc0120

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestMinimumTotal(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `[[2],[3,4],[6,5,7],[4,1,8,3]]`,
			want:  11,
		},
		{
			input: `[[-10]]`,
			want:  -10,
		},
		{
			input: `[[-1],[2,3],[1,-1,-3]]`,
			want:  -1,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			triangle, err := readTriangle(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			got := minimumTotal(triangle)
			if tt.want != got {
				t.Fatalf("expect %d, got %d", tt.want, got)
			}
		})
	}
}

func readTriangle(input string) ([][]int, error) {
	input = strings.TrimSpace(input)
	if !strings.HasPrefix(input, "[[") || !strings.HasSuffix(input, "]]") {
		return nil, fmt.Errorf("input should start with [[ and end with ]]")
	}
	input = strings.TrimPrefix(input, "[[")
	input = strings.TrimSuffix(input, "]]")
	input = strings.ReplaceAll(input, " ", "")
	rows := strings.Split(input, "],[")
	triangle := make([][]int, len(rows))
	for i, row := range rows {
		cols := strings.Split(row, ",")
		triangle[i] = make([]int, len(cols))
		for j, col := range cols {
			v, err := strconv.Atoi(col)
			if err != nil {
				return nil, err
			}
			triangle[i][j] = v
		}
	}
	return triangle, nil
}
