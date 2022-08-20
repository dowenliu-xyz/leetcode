package main

func twoSum(nums []int, target int) []int {
	//return solution1(nums, target)
	return solution2(nums, target)
}

func solution1(nums []int, target int) []int {
	// 暴力枚举 注意返回数组下标
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func solution2(nums []int, target int) []int {
	// 查表法
	table := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if j, ok := table[target-nums[i]]; ok {
			return []int{j, i}
		}
		table[nums[i]] = i
	}
	return nil
}
