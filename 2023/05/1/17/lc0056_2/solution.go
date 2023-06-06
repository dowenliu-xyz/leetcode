package lc0056_2

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var ans [][]int
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
