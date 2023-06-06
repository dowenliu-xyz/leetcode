package lc0455_4

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	si, gi := len(s)-1, len(g)-1
	ans := 0
	for si >= 0 && gi >= 0 {
		if s[si] >= g[gi] {
			si--
			ans++
		}
		gi--
	}
	return ans
}
