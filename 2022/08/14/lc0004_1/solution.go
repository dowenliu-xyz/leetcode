package main

import (
	"math"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//return solution1(nums1, nums2)
	//return solution2(nums1, nums2)
	//return solution3(nums1, nums2)
	//return solution4(nums1, nums2)
	return solution5(nums1, nums2)
}

func solution1(nums1, nums2 []int) float64 {
	// 合并后排序，直接取中位数。
	// 时间复杂度 O(n*log(n))
	nums := make([]int, 0, len(nums1)+len(nums2))
	nums = append(nums, nums1...)
	nums = append(nums, nums2...)
	n := len(nums)
	if n == 0 {
		return 0.0
	}
	sort.Ints(nums)
	if n&1 == 1 {
		return float64(nums[n/2])
	} else {
		num1, num2 := nums[n/2-1], nums[n/2]
		return float64(num1+num2) / 2
	}
}

func solution2(nums1, nums2 []int) float64 {
	// 合并有序数组
	// 时间复杂度 O(m+n)
	nums := make([]int, 0, len(nums1)+len(nums2))
	p1, p2 := 0, 0
	for p1 < len(nums1) || p2 < len(nums2) {
		if p1 == len(nums1) {
			nums = append(nums, nums2[p2])
			p2++
			continue
		}
		if p2 == len(nums2) {
			nums = append(nums, nums1[p1])
			p1++
			continue
		}
		if nums1[p1] < nums2[p2] {
			nums = append(nums, nums1[p1])
			p1++
			continue
		}
		nums = append(nums, nums2[p2])
		p2++
	}
	n := len(nums)
	if n == 0 {
		return 0.0
	}
	if n&1 == 1 {
		return float64(nums[n/2])
	} else {
		num1, num2 := nums[n/2-1], nums[n/2]
		return float64(num1+num2) / 2
	}
}

func solution3(nums1, nums2 []int) float64 {
	// 不合并，直接查找
	// 时间复杂度 O(m+n)
	nums := len(nums1) + len(nums2)
	var pre, cur int
	p1, p2, counter, mp := 0, 0, 0, nums/2
	for p1 < len(nums1) || p2 < len(nums2) {
		pre = cur
		if p1 == len(nums1) {
			cur = nums2[p2]
			p2++
		} else if p2 == len(nums2) {
			cur = nums1[p1]
			p1++
		} else if nums1[p1] < nums2[p2] {
			cur = nums1[p1]
			p1++
		} else {
			cur = nums2[p2]
			p2++
		}
		if counter == mp {
			break
		}
		counter++
	}
	if nums&1 == 1 {
		return float64(cur)
	} else {
		return (float64(cur + pre)) / 2
	}
}

func solution4(nums1, nums2 []int) float64 {
	// 求第k小值
	// 时间复杂度 O(log(m + n)
	totalLength := len(nums1) + len(nums2)
	if totalLength == 0 {
		return 0
	}
	if totalLength&1 == 1 {
		return float64(getKthElement(nums1, nums2, totalLength/2+1))
	} else {
		return float64(getKthElement(nums1, nums2, totalLength/2)+getKthElement(nums1, nums2, totalLength/2+1)) / 2
	}
}

func getKthElement(nums1, nums2 []int, k int) int {
	p1, p2 := 0, 0
	for {
		if p1 == len(nums1) {
			return nums2[p2+k-1]
		}
		if p2 == len(nums2) {
			return nums1[p1+k-1]
		}
		if k == 1 {
			return min(nums1[p1], nums2[p2])
		}
		hafl := k / 2
		newP1 := min(p1+hafl, len(nums1)) - 1
		newP2 := min(p2+hafl, len(nums2)) - 1
		if nums1[newP1] < nums2[newP2] {
			k -= newP1 - p1 + 1
			p1 = newP1 + 1
		} else {
			k -= newP2 - p2 + 1
			p2 = newP2 + 1
		}
	}
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func solution5(nums1, nums2 []int) float64 {
	// 划分数组
	// 时间复杂度 O(log(min(m, n)))
	m, n := len(nums1), len(nums2)
	if m+n == 0 {
		return 0
	}
	if m > n {
		return solution5(nums2, nums1)
	}
	mp := m + (n-m+1)/2
	// m 必定小于等于 mp，因为 m <= n。同理 n 必定大于等于 mp
	p1, p2 := m-1, mp-m-1
	// 找 left1 严格小于 right2 的位置
	for p1 >= 0 && nums1[p1] >= nums2[p2+1] {
		p1, p2 = p1-1, p2+1
	}
	left1, left2, right1, right2 := math.MinInt, math.MinInt, math.MaxInt, math.MaxInt
	if p1 >= 0 {
		// 存在 left1
		left1 = nums1[p1]
	}
	if p1+1 < m {
		// 存在 right1
		right1 = nums1[p1+1]
	}
	if p2 >= 0 {
		// 存在 left2
		left2 = nums2[p2]
	}
	if p2+1 < n {
		// 存在 right2
		right2 = nums2[p2+1]
	}
	if (m+n)&1 == 1 {
		return float64(max(left1, left2))
	} else {
		return float64(max(left1, left2)+min(right1, right2)) / 2
	}
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
