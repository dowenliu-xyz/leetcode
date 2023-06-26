package lc0435

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestEraseOverlapIntervals(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `[[1,2],[2,3],[3,4],[1,3]]`,
			want:  1,
		},
		{
			input: `[ [1,2], [1,2], [1,2] ]`,
			want:  2,
		},
		{
			input: `[ [1,2], [2,3] ]`,
			want:  0,
		},
		{
			input: `[[1,100],[11,22],[1,11],[2,12]]`,
			want:  2,
		},
		{
			input: `[[-52,31],[-73,-26],[82,97],[-65,-11],[-62,-49],[95,99],[58,95],[-31,49],[66,98],[-63,2],[30,47],[-40,-26]]`,
			want:  7,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			intervals := parseInput(t, tt.input)
			got := eraseOverlapIntervals(intervals)
			if got != tt.want {
				t.Errorf("expect %d, got %d", tt.want, got)
			}
		})
	}
}

func parseInput(t *testing.T, input string) [][]int {
	t.Helper()

	input = strings.ReplaceAll(input, " ", "")
	input = input[2 : len(input)-2] // remove '[[' and ']]'
	intervalStrs := strings.Split(input, "],[")
	intervals := make([][]int, 0, len(intervalStrs))
	for _, str := range intervalStrs {
		str = strings.TrimSpace(str)
		ends := strings.Split(str, ",")
		if len(ends) != 2 {
			t.Fatalf("invalid interval: [%s]", str)
		}
		start, err := strconv.ParseInt(ends[0], 10, 64)
		if err != nil {
			t.Fatal(err)
		}
		end, err := strconv.ParseInt(ends[1], 10, 64)
		if err != nil {
			t.Fatal(err)
		}
		intervals = append(intervals, []int{int(start), int(end)})
	}
	return intervals
}
