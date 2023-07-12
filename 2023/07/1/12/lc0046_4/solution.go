package lc0046

func permute(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return nil
	}
	var ans [][]int
	used := make([]bool, n)
	var dfs func(p []int)
	dfs = func(p []int) {
		if len(p) == n {
			tmp := make([]int, n)
			copy(tmp, p)
			ans = append(ans, tmp)
			return
		}
		for i, num := range nums {
			if used[i] {
				continue
			}
			used[i] = true
			dfs(append(p, num))
			used[i] = false
		}
	}
	dfs(make([]int, 0, n))
	return ans
}
