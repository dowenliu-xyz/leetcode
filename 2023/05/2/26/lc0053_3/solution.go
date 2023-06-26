package lc0053

import "math"

func maxSubArray(nums []int) int {
	//return solution1(nums)
	return solution2(nums)
}

func solution1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans, sum := math.MinInt, 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		ans = max(ans, sum)
		if sum < 0 {
			sum = 0
		}
	}
	return ans
}

func solution2(nums []int) int {
	return get(nums, 0, len(nums)-1).maxSubSum
}

type Status struct {
	leftMaxSum  int
	rightMaxSum int
	totalSum    int
	maxSubSum   int
}

func get(nums []int, left, right int) Status {
	if left == right {
		return Status{
			leftMaxSum:  nums[left],
			rightMaxSum: nums[left],
			totalSum:    nums[left],
			maxSubSum:   nums[left],
		}
	}
	m := left + (right-left)>>1
	s1, s2 := get(nums, left, m), get(nums, m+1, right)
	return pushUp(s1, s2)
}

func pushUp(left, right Status) Status {
	return Status{
		leftMaxSum:  max(left.leftMaxSum, left.totalSum+right.leftMaxSum),
		rightMaxSum: max(right.rightMaxSum, left.rightMaxSum+right.totalSum),
		totalSum:    left.totalSum + right.totalSum,
		maxSubSum:   max(max(left.maxSubSum, right.maxSubSum), left.rightMaxSum+right.leftMaxSum),
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
