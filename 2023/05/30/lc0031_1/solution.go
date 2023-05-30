package lc0031_1

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	i, j := len(nums)-2, len(nums)-1
	for i >= 0 {
		if nums[i] < nums[i+1] {
			break
		}
		i--
	}
	if i >= 0 {
		for j >= i {
			if nums[j] > nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
			j--
		}
	}
	for l, r := i+1, len(nums)-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}
