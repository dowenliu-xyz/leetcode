package main

func findRepeatNumber(nums []int) int {
	//return solution1(nums)
	return solution2(nums)
}

func solution1(nums []int) int {
	// 查表
	seen := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if seen[nums[i]] {
			return nums[i]
		}
		seen[nums[i]] = true
	}
	return -1
}

func solution2(nums []int) int {
	i := 0
	for i < len(nums) {
		if nums[i] == i {
			i++
			continue
		}
		if nums[i] == nums[nums[i]] {
			return nums[i]
		}
		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
	}
	return -1
}
