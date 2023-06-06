package lc0300_1

func lengthOfLIS(nums []int) int {
	mem := make([]int, 0, len(nums))
	for _, num := range nums {
		if len(mem) > 0 && mem[len(mem)-1] >= num {
			l, r := 0, len(mem)
			for l <= r {
				m := l + (r-l)>>1
				if mem[m] < num {
					l = m + 1
				} else {
					r = m - 1
				}
			}
			mem[l] = num
		} else {
			mem = append(mem, num)
		}
	}
	return len(mem)
}
