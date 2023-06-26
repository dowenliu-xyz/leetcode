package lc1480

func runningSum(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}
	ans := make([]int, n)
	ans[0] = nums[0]
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] + nums[i]
	}
	return ans
}
