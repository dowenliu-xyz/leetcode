package lc0239

func maxSlidingWindow(nums []int, k int) []int {
	var win []int
	var ans []int
	for i := 0; i < len(nums); i++ {
		for len(win) > 0 && nums[win[len(win)-1]] <= nums[i] {
			win = win[:len(win)-1]
		}
		win = append(win, i)
		if i < k-1 {
			continue
		}
		if win[0] <= i-k {
			win = win[1:]
		}
		ans = append(ans, nums[win[0]])
	}
	return ans
}
