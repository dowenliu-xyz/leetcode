package lc0128

func longestConsecutive(nums []int) int {
	set := map[int]bool{}
	for _, num := range nums {
		set[num] = true
	}
	ans := 0
	for _, num := range nums {
		if set[num-1] {
			continue
		}
		l := 1
		for set[num+1] {
			l++
			num++
		}
		if l > ans {
			ans = l
		}
	}
	return ans
}
