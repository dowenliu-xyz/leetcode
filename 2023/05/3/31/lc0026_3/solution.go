package lc0026_3

func removeDuplicates(nums []int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[k] == nums[i] {
			continue
		}
		k++
		nums[i], nums[k] = nums[k], nums[i]
	}
	return k + 1
}
