package lc0004_1

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	if totalLen == 0 {
		return 0.0
	}
	if totalLen&1 == 1 {
		return float64(getKthElement(nums1, nums2, totalLen>>1+1))
	}
	l, r := getKthElement(nums1, nums2, totalLen>>1), getKthElement(nums1, nums2, totalLen>>1+1)
	return float64(l) + float64(r-l)/2
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k >> 1
		newIndex1, newIndex2 := min(index1+half-1, len(nums1)-1), min(index2+half-1, len(nums2)-1)
		if nums1[newIndex1] <= nums2[newIndex2] {
			k -= newIndex1 - index1 + 1
			index1 = newIndex1 + 1
		} else {
			k -= newIndex2 - index2 + 1
			index2 = newIndex2 + 1
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
