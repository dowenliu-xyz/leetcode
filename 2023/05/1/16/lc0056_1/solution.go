package lc0056_1

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n < 2 {
		return intervals
	}
	var ans [][]int
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := range intervals {
		if len(ans) == 0 || ans[len(ans)-1][1] < intervals[i][0] {
			ans = append(ans, intervals[i])
			continue
		}
		ans[len(ans)-1][1] = max(ans[len(ans)-1][1], intervals[i][1])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
