package lc0455_3

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	ans := 0

	for i, j := len(g)-1, len(s)-1; j >= 0 && i >= 0; i-- {
		if g[i] <= s[j] {
			ans++
			j--
		}
	}
	return ans
}
