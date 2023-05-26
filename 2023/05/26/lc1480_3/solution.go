package lc1480_3

func runningSum(nums []int) []int {
	ans := make([]int, len(nums))
	copy(ans, nums)
	for i := 1; i < len(nums); i++ {
		ans[i] += ans[i-1]
	}
	return ans
}
