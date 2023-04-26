package of0051_1

func reversePairs(nums []int) int {
	// 归并排序
	n := len(nums)
	if n <= 1 {
		return 0
	}
	m := n >> 1
	ans := reversePairs(nums[:m]) + reversePairs(nums[m:])
	tmp := make([]int, 0, n)
	i, j := 0, m
	for i < m && j < n {
		if nums[i] > nums[j] {
			tmp = append(tmp, nums[j])
			ans += m - i
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
		tmp = append(tmp, nums[j:n]...)
	}
	copy(nums, tmp)
	return ans
}
