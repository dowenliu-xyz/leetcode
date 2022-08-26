package main

import "sort"

func containsDuplicate(nums []int) bool {
	//return solution1(nums)
	//return solution2(nums)
	return solution3(nums)
}

func solution1(nums []int) bool {
	// 暴力枚举
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

func solution2(nums []int) bool {
	// 排序再检查相临是否有相同
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

func solution3(nums []int) bool {
	// hash 排重
	seen := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if seen[nums[i]] {
			return true
		}
		seen[nums[i]] = true
	}
	return false
}
