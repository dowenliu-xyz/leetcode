package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	l, r := 0, 1
	for r < len(nums) {
		if nums[l] == nums[r] {
			r++
			continue
		}
		l++
		nums[l], nums[r] = nums[r], nums[l]
		r++
	}
	return l + 1
}
