package lc0056

import "sort"

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := make([][]int, 0, n)
	ans = append(ans, intervals[0])
	for _, interval := range intervals[1:] {
		last := ans[len(ans)-1]
		if last[1] >= interval[0] {
			last[1] = max(last[1], interval[1])
			continue
		}
		ans = append(ans, interval)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
