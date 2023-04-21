package lc0455_1

import "sort"

func findContentChildren(g []int, s []int) int {
	//return solution1(g, s)
	return solution2(g, s)
}

func solution2(g, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var count int
	for i, j := len(g)-1, len(s)-1; i >= 0; i-- {
		if j < 0 {
			break
		}
		if g[i] > s[j] {
			continue
		}
		count++
		j--
	}
	return count
}

func solution1(g []int, s []int) int {
	sort.Ints(s)              // 饼干尺寸从小到大排序
	t := make([]bool, len(s)) // t[i] == true 表示 s[i] 已经被分配了
	sort.Ints(g)              // 小孩胃口从小到大排序
	var count int             // 分配的饼干数量
	for _, gi := range g {    // 从胃口最小的小孩开始
		for i, si := range s { // 从尺寸最小的饼干开始
			if !t[i] && si >= gi { // 如果饼干尺寸大于等于小孩胃口
				t[i] = true // 标记饼干已经被分配
				count++     // 分配的饼干数量加一
				break       // 退出循环，分配下一个小孩
			}
		}
	}
	return count // 返回分配的饼干数量
}
