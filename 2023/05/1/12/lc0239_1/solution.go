package lc0239_1

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n < 1 {
		return nil
	}
	ans := make([]int, 0, n-k)
	var q []int // q 为双向队列，其中存储下标，对应值单调递增
	for i := 0; i < n; i++ {
		for len(q) > 0 && nums[q[len(q)-1]] < nums[i] { // 清理尾部，以保证入队后，队列单调递增
			q = q[:len(q)-1]
		}
		q = append(q, i)                // i 入队
		for len(q) > 0 && q[0] <= i-k { // 超出左边界的元素出队列
			q = q[1:]
		}
		if i < k-1 { // 窗口还没填满
			continue
		}
		ans = append(ans, nums[q[0]]) // 窗口队列头部总是当前窗口内的最大值
	}
	return ans
}
