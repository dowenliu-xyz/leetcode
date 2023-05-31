package lc0031_2

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
		for j > i {
			if nums[j] > nums[i] {
				break
			}
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	for l, r := i+1, n-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}
