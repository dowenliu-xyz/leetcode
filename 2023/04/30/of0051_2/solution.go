package of0051_2

func reversePairs(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	m := n / 2
	ans := reversePairs(nums[:m]) + reversePairs(nums[m:])
	i, j := 0, m
	tmp := make([]int, 0, n)
	for i < m && j < n {
		if nums[i] > nums[j] {
			tmp = append(tmp, nums[j])
			j++
			ans += m - i
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
