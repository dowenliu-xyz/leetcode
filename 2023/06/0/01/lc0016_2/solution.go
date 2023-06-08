package lc0016

import "sort"

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	sort.Ints(nums)
	ans := nums[0] + nums[1] + nums[2]
	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		second, third := first+1, n-1
		for second < third {
			sum := nums[first] + nums[second] + nums[third]
			if abs(target-ans) > abs(target-sum) {
				ans = sum
			}
			if sum == target {
				return sum
			}
			if sum < target {
				for ; second < third && nums[second] == nums[second+1]; second++ {
				}
				second++
			} else {
				for ; second < third && nums[third] == nums[third-1]; third-- {
				}
				third--
			}
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
