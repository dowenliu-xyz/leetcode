package main

func runningSum(nums []int) []int {
	return solution2(nums)
}

func solution1(nums []int) []int {
	// 不修改原数据
	if len(nums) == 0 {
		return nil
	}
	ans := make([]int, len(nums))
	ans[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i-1] + nums[i]
	}
	return ans
}

func solution2(nums []int) []int {
	// 原地修改
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]
	}
	return nums
}
