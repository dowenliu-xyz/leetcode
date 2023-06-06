package lc0056_3

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var ans [][]int
	for _, interval := range intervals {
		if len(ans) > 0 && ans[len(ans)-1][1] >= interval[0] {
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], interval[1])
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
