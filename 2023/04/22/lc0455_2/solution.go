package lc0455_2

import "sort"

func findContentChildren(g, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	res := 0
	for i, j := len(g)-1, len(s)-1; j >= 0; {
		if i < 0 { // 所有小孩的期望已经判断完
			break
		}
		if g[i] <= s[j] {
			j--   // 饼干已分配
			res++ // 更新满足小孩数
		}
		i-- // 下个小孩
	}
	return res
}
