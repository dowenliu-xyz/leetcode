package lc0239

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n < k {
		return nil
	}
	win := make([]int, 0, k)
	ans := make([]int, 0, n-k)
	for i, num := range nums {
		for len(win) > 0 && nums[win[len(win)-1]] <= num {
			win = win[:len(win)-1]
		}
		win = append(win, i)
		if i < k-1 {
			continue
		}
		for len(win) > 0 && win[0] < i-k+1 {
			win = win[1:]
		}
		ans = append(ans, nums[win[0]])
	}
	return ans
}
