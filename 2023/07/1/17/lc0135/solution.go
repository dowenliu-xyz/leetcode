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
	rignt := 1
	ans := max(left[n-1], rignt)
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			rignt++
		} else {
			rignt = 1
		}
		ans += max(left[i], rignt)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
