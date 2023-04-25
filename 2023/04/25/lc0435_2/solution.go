package lc0435_2

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	// 处理特殊值
	if n <= 0 {
		// 无需移除任何区间
		return 0
	}
	// 按区间右边界排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	ans, right := 1, intervals[0][1] // 最多保留区间数，最后保持区间的右边界
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < right {
			// 当前区间与最后保持区间有重叠，移除
			continue
		}
		ans++                   // 更新保留数
		right = intervals[i][1] // 变更最后保留区间右边界
	}
	return n - ans // 所求为最少移除数。区间总数 - 最大保留数
}
