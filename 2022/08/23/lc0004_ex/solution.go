package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//return solution1(nums1, nums2)
	return solution2(nums1, nums2)
}

func solution1(nums1, nums2 []int) float64 {
	// 原地数到中位数 练习
	m, n := len(nums1), len(nums2)
	if m+n == 0 {
		return 0
	}
	i1, i2, mid := 0, 0, (m+n)>>1
	var pre, cur int
	for i1+i2 <= mid && (i1 < m || i2 < n) {
		pre = cur
		if i1 == m {
			cur = nums2[i2]
			i2++
		} else if i2 == n {
			cur = nums1[i1]
			i1++
		} else if nums1[i1] < nums2[i2] {
			cur = nums1[i1]
			i1++
		} else {
			cur = nums2[i2]
			i2++
		}
	}
	if (m+n)&1 == 1 {
		return float64(cur)
	}
	return float64(pre) + float64(cur-pre)/2
}

func solution2(nums1, nums2 []int) float64 {
	// 转化为求第K小值问题 练习
	m, n := len(nums1), len(nums2)
	totalLength := m + n
	if totalLength == 0 {
		return 0
	}
	if m > n {
		return solution2(nums2, nums1)
	}
	if totalLength&1 == 1 {
		return float64(getKthElement(nums1, nums2, totalLength>>1+1))
	}
	mid1, mid2 := getKthElement(nums1, nums2, totalLength>>1), getKthElement(nums1, nums2, totalLength>>1+1)
	return float64(mid1) + float64(mid2-mid1)/2
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
		newOffset1 := min(offset1+half-1, len(nums1)-1)
		newOffset2 := min(offset2+half-1, len(nums2)-1)
		if nums1[newOffset1] < nums2[newOffset2] {
			newOffset1 += 1
			k -= newOffset1 - offset1
			offset1 = newOffset1
		} else {
			newOffset2 += 1
			k -= newOffset2 - offset2
			offset2 = newOffset2
		}
	}
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
