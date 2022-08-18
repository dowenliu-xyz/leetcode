package main

import "sort"

func containsDuplicate(nums []int) bool {
	//return solution1(nums)
	return solution2(nums)
}

func solution1(nums []int) bool {
	seen := map[int]bool{}
	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}

func solution2(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}
