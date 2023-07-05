package lc0239

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n < k {
		return nil
	}
	var q []int
	ans := make([]int, 0, n-k)
	for i, num := range nums {
		for len(q) > 0 && nums[q[len(q)-1]] <= num {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		if i < k-1 {
			continue
		}
		for len(q) > 0 && q[0]+k-1 < i {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}
