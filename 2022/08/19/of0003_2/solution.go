package main

func findRepeatNumber(nums []int) int {
	//return solution1(nums)
	return solution2(nums)
}

func solution1(nums []int) int {
	// hash 排重
	seen := map[int]bool{}
	for _, num := range nums {
		if seen[num] {
			return num
		}
		seen[num] = true
	}
	return -1
}

func solution2(nums []int) int {
	// 按照题目条件如果没有重复则所有值与位置下标刚好可一一对应
	i := 0
	for i < len(nums) {
		if nums[i] == i {
			i++
			continue
		}
		if nums[nums[i]] == nums[i] {
			return nums[i]
		}
		nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
	}
	return -1
}
