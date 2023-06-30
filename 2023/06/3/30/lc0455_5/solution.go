package lc0455

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	count, i, j := 0, len(g)-1, len(s)-1
	for i >= 0 && j >= 0 { // 当有小孩未被询问且手上还有饼干时
		if g[i] <= s[j] {
			// 可以满足小孩 i
			count++
			i--
			j--
			continue
		}
		// 看是否能满足小孩 i-1
		i--
	}
	return count
}
