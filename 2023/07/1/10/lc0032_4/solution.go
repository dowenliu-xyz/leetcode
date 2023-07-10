package lc0031

func nextPermutation(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}
	i := n - 2
	for i >= 0 {
		if nums[i] < nums[i+1] {
			break
		}
		i--
	}
	if i >= 0 {
		j := n - 1
		for i < j {
			if nums[i] < nums[j] {
				break
			}
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	l, r := i+1, n-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
