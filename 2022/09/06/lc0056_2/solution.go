package main

import "sort"

func merge(intervals [][]int) [][]int {
	merged := make([][]int, 0, len(intervals))
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 0; i < len(intervals); i++ {
		if len(merged) == 0 || merged[len(merged)-1][1] < intervals[i][0] {
			merged = append(merged, intervals[i])
			continue
		}
		merged[len(merged)-1][1] = max(merged[len(merged)-1][1], intervals[i][1])
	}
	return merged
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
