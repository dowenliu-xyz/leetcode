package main

func search(nums []int, target int) int {
	// 注意返回下标
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left+1)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			// nums[mid] < target
			left = mid + 1
		}
	}
	return -1
}
