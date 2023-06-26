package lc0053

func maxSubArray(nums []int) int {
	//return solution1(nums)
	return solution2(nums)
}

func solution1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans, sum := nums[0], 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		if sum > ans {
			ans = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return ans
}

type Status struct {
	LeftMaxSum  int
	RightMaxSum int
	MaxSubSum   int
	TotalSum    int
}

func solution2(nums []int) int {
	return get(nums, 0, len(nums)-1).MaxSubSum
}

func get(nums []int, left, right int) Status {
	if left == right {
		return Status{
			LeftMaxSum:  nums[left],
			RightMaxSum: nums[left],
			MaxSubSum:   nums[left],
			TotalSum:    nums[left],
		}
	}
	m := left + (right-left)>>1
	s1, s2 := get(nums, left, m), get(nums, m+1, right)
	return pushUp(s1, s2)
}

func pushUp(left, right Status) Status {
	return Status{
		LeftMaxSum:  max(left.LeftMaxSum, left.TotalSum+right.LeftMaxSum),
		RightMaxSum: max(right.RightMaxSum, right.TotalSum+left.RightMaxSum),
		MaxSubSum:   max(left.RightMaxSum+right.LeftMaxSum, max(left.MaxSubSum, right.MaxSubSum)),
		TotalSum:    left.TotalSum + right.TotalSum,
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
