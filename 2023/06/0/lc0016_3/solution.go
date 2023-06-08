package lc0016

import "sort"

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	sort.Ints(nums)
	ans := nums[0] + nums[1] + nums[2]
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if abs(ans-target) > abs(sum-target) {
				ans = sum
			}
			if sum == target {
				return ans
			}
			if sum < target {
				for ; l < r && nums[l] == nums[l+1]; l++ {
				}
				l++
			} else {
				for ; l < r && nums[r] == nums[r-1]; r-- {
				}
				r--
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
