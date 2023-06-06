package lc0015_3

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	for first := 0; first < len(nums); first++ {
		if nums[first] > 0 {
			break
		}
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		second, third := first+1, len(nums)-1
		for second < third {
			sum := nums[first] + nums[second] + nums[third]
			if sum == 0 {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
				for ; second < third && nums[second] == nums[second+1]; second++ {
				}
				for ; second < third && nums[third-1] == nums[third]; third-- {
				}
				second++
				third--
			} else if sum < 0 {
				second++
			} else {
				third--
			}
		}
	}
	return ans
}
