package lc0004

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	if total&1 == 1 {
		return float64(findKthElement(nums1, nums2, total>>1+1))
	}
	l, r := findKthElement(nums1, nums2, total>>1), findKthElement(nums1, nums2, total>>1+1)
	return float64(l) + float64(r-l)/2
}

func findKthElement(nums1, nums2 []int, k int) int {
	offset1, offset2 := 0, 0
	for {
		if offset1 == len(nums1) {
			return nums2[offset2+k-1]
		}
		if offset2 == len(nums2) {
			return nums1[offset1+k-1]
		}
		if k == 1 {
			return min(nums1[offset1], nums2[offset2])
		}
		half := k >> 1
		newOffset1, newOffset2 := min(len(nums1), offset1+half)-1, min(len(nums2), offset2+half)-1
		if nums1[newOffset1] <= nums2[newOffset2] {
			k -= newOffset1 + 1 - offset1
			offset1 = newOffset1 + 1
		} else {
			k -= newOffset2 + 1 - offset2
			offset2 = newOffset2 + 1
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
