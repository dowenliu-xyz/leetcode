package lc0026

func removeDuplicates(nums []int) int {
	k, n := 0, len(nums)
	if n < 2 {
		return n
	}
	for i := 0; i < n; i++ {
		if nums[k] == nums[i] {
			continue
		}
		k++
		nums[k], nums[i] = nums[i], nums[k]
	}
	return k + 1
}
