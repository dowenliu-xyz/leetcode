package of0051_3

func reversePairs(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	m := n >> 1
	ans := reversePairs(nums[:m]) + reversePairs(nums[m:])
	tmp, i, j := make([]int, 0, n), 0, m
	for i < m && j < n {
		if nums[i] > nums[j] {
			ans += m - i
			tmp = append(tmp, nums[j])
			j++
		} else {
			tmp = append(tmp, nums[i])
			i++
		}
	}
	if i < m {
		tmp = append(tmp, nums[i:m]...)
	}
	if j < n {
		tmp = append(tmp, nums[j:]...)
	}
	copy(nums, tmp)
	return ans
}
