package lc0088

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k], i, k = nums1[i], i-1, k-1
		} else {
			nums1[k], j, k = nums2[j], j-1, k-1
		}
	}
	if j >= 0 {
		copy(nums1[:k+1], nums2[:j+1])
	}
}
