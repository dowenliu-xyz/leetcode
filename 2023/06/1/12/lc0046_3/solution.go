package lc0046

func permute(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return nil
	}
	var ans [][]int
	mem := make([]bool, n)
	var search func(p []int)
	search = func(p []int) {
		if len(p) == n {
			tmp := make([]int, n)
			copy(tmp, p)
			ans = append(ans, tmp)
			return
		}
		for i := 0; i < n; i++ {
			if mem[i] {
				continue
			}
			mem[i] = true
			search(append(p, nums[i]))
			mem[i] = false
		}
	}
	search(make([]int, 0, n))
	return ans
}
