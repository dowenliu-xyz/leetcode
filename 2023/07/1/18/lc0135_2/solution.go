package lc0135

func candy(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}
	left := make([]int, n)
	left[0] = 1
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	right, ans := 1, max(1, left[n-1])
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		ans += max(left[i], right)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
