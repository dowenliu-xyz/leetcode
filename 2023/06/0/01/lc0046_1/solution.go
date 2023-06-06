package lc0046_1

func permute(nums []int) [][]int {
	var ans [][]int
	n := len(nums)
	mem := make([]bool, n)
	var search func(result []int)
	search = func(result []int) {
		if len(result) == n {
			tmp := make([]int, n)
			copy(tmp, result)
			ans = append(ans, tmp)
			return
		}
		for i := 0; i < n; i++ {
			if mem[i] {
				continue
			}
			mem[i] = true
			search(append(result, nums[i]))
			mem[i] = false
		}
	}
	search(make([]int, 0, n))
	return ans
}
