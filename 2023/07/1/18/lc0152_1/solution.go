package lc0152

func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	maxF, minF := make([]int, n), make([]int, n)
	copy(maxF, nums)
	copy(minF, nums)
	for i := 1; i < n; i++ {
		maxF[i] = max(maxF[i-1]*nums[i], max(nums[i], minF[i-1]*nums[i]))
		minF[i] = min(minF[i-1]*nums[i], min(nums[i], maxF[i-1]*nums[i]))
	}
	ans := maxF[0]
	for i := 1; i < n; i++ {
		ans = max(ans, maxF[i])
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
