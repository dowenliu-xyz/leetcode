package lc1480_1

func runningSum(nums []int) []int {
	ans := make([]int, len(nums))
	copy(ans, nums)
	for i := 1; i < len(ans); i++ {
		ans[i] += ans[i-1]
	}
	return ans
}
