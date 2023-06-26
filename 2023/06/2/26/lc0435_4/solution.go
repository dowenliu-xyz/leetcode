package lc0435

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n < 2 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1] // 按右边界升序排序
	})
	ans, right := 0, intervals[0][1]
	for i := 1; i < n; i++ {
		if intervals[i][0] < right {
			ans++
		} else {
			right = intervals[i][1]
		}
	}
	return ans
}
