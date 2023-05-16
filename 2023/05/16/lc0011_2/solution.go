package lc0011_2

func maxArea(height []int) int {
	n := len(height)
	if n < 2 {
		return 0
	}
	ans, l, r := 0, 0, n-1
	for l < r {
		ans = max(ans, (r-l)*min(height[l], height[r]))
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
