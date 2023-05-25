package lc0004_2

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength == 0 {
		return 0
	}
	if totalLength&1 == 1 {
		return float64(getKthElement(nums1, nums2, totalLength>>1+1))
	}
	m1, m2 := getKthElement(nums1, nums2, totalLength>>1), getKthElement(nums1, nums2, totalLength>>1+1)
	return float64(m1) + float64(m2-m1)/2
}

func getKthElement(nums1, nums2 []int, k int) int {
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
		newOffset1, newOffset2 := min(offset1+half-1, len(nums1)-1), min(offset2+half-1, len(nums2)-1)
		if nums1[newOffset1] < nums2[newOffset2] {
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
