package main

import (
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//return solution1(nums1, nums2)
	//return solution2(nums1, nums2)
	//return solution3(nums1, nums2)
	return solution4(nums1, nums2)
}

func solution1(nums1, nums2 []int) float64 {
	// 合并排序再计算中位数
	totalLength := len(nums1) + len(nums2)
	nums := make([]int, 0, totalLength)
	nums = append(nums, nums1...)
	nums = append(nums, nums2...)
	if totalLength == 0 {
		return 0
	}
	sort.Ints(nums)
	if totalLength&1 == 1 {
		return float64(nums[totalLength>>1])
	} else {
		mid1, mid2 := totalLength>>1-1, totalLength>>1
		return float64(nums[mid1]) + float64(nums[mid2]-nums[mid1])/2
	}
}

func solution2(nums1, nums2 []int) float64 {
	// 合并有序数组，再计算中位数
	totalLength := len(nums1) + len(nums2)
	if totalLength == 0 {
		return 0
	}
	nums := make([]int, 0, totalLength)
	i1, i2 := 0, 0
	for i1 < len(nums1) || i2 < len(nums2) {
		if i1 == len(nums1) {
			nums = append(nums, nums2[i2])
			i2++
			continue
		}
		if i2 == len(nums2) {
			nums = append(nums, nums1[i1])
			i1++
			continue
		}
		if nums1[i1] < nums2[i2] {
			nums = append(nums, nums1[i1])
			i1++
		} else {
			nums = append(nums, nums2[i2])
			i2++
		}
	}
	if totalLength&1 == 1 {
		return float64(nums[totalLength>>1])
	} else {
		mid1, mid2 := totalLength>>1-1, totalLength>>1
		return float64(nums[mid1]) + float64(nums[mid2]-nums[mid1])/2
	}
}

func solution3(nums1, nums2 []int) float64 {
	// 不合并数据，直接数到中位数
	n, m := len(nums1), len(nums2)
	totalLength := n + m
	if totalLength == 0 {
		return 0
	}
	i1, i2, count, mid := 0, 0, -1, totalLength>>1
	var preNum, curNum int
	for count < mid && (i1 < n || i2 < m) {
		preNum = curNum
		if i1 == n {
			curNum = nums2[i2]
			i2++
		} else if i2 == m {
			curNum = nums1[i1]
			i1++
		} else if nums1[i1] < nums2[i2] {
			curNum = nums1[i1]
			i1++
		} else {
			curNum = nums2[i2]
			i2++
		}
		count++
	}
	if totalLength&1 == 1 {
		return float64(curNum)
	}
	return float64(preNum) + float64(curNum-preNum)/2
}

func solution4(nums1, nums2 []int) float64 {
	// 转化为求第K小值问题
	totalLength := len(nums1) + len(nums2)
	if totalLength == 0 {
		return 0
	}
	if totalLength&1 == 1 {
		// 注意：totalLength>>1 是中位数全局下标，再+1才是序号
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
			k -= newOffset1 - offset1 + 1
			offset1 = newOffset1 + 1
		} else {
			k -= newOffset2 - offset2 + 1
			offset2 = newOffset2 + 1
		}
	}
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
