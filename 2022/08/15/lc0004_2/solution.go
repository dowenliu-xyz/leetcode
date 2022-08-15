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
	// 合并，排序，取中位数
	// O(n*log(n)) 不合题意
	totalLength := len(nums1) + len(nums2)
	if totalLength == 0 {
		return 0.0
	}
	nums := make([]int, 0, totalLength)
	nums = append(nums, nums1...)
	nums = append(nums, nums2...)
	sort.Ints(nums)
	if totalLength&1 == 1 {
		return float64(nums[totalLength/2])
	}
	mp1, mp2 := totalLength/2-1, totalLength/2
	return float64(nums[mp1]) + (float64(nums[mp2]-nums[mp1]))/2.0
}

func solution2(nums1, nums2 []int) float64 {
	// 合并有序数据，再取中位数
	// O(m+n) 不合题意
	m, n := len(nums1), len(nums2)
	if m+n == 0 {
		return 0
	}
	nums, p1, p2 := make([]int, 0, m+n), 0, 0
	for p1 < m || p2 < n {
		if p1 == m {
			nums = append(nums, nums2[p2])
			p2++
		} else if p2 == n {
			nums = append(nums, nums1[p1])
			p1++
		} else if nums1[p1] < nums2[p2] {
			nums = append(nums, nums1[p1])
			p1++
		} else {
			nums = append(nums, nums2[p2])
			p2++
		}
	}
	mp := (m + n) / 2
	if (m+n)&1 == 1 {
		return float64(nums[mp])
	}
	return float64(nums[mp-1]) + float64(nums[mp]-nums[mp-1])/2.0
}

func solution3(nums1, nums2 []int) float64 {
	// 不合并直接查找
	// O(m+n) 不合题意
	m, n := len(nums1), len(nums2)
	totalLength := m + n
	if totalLength == 0 {
		return 0
	}
	p1, p2, count, mp := 0, 0, 0, totalLength/2
	pre, cur := 0, 0
	for p1 < m || p2 < n {
		pre = cur
		if p1 == m {
			cur = nums2[p2]
			p2++
		} else if p2 == n {
			cur = nums1[p1]
			p1++
		} else if nums1[p1] < nums2[p2] {
			cur = nums1[p1]
			p1++
		} else {
			cur = nums2[p2]
			p2++
		}
		if count == mp {
			break
		}
		count++
	}
	if totalLength&1 == 1 {
		return float64(cur)
	} else {
		return float64(pre) + float64(cur-pre)/2
	}
}

func solution4(nums1, nums2 []int) float64 {
	// 求第K小值
	// O(log(m+n))
	totalLength := len(nums1) + len(nums2)
	if totalLength == 0 {
		return 0
	}
	if totalLength&1 == 1 {
		return float64(getKthElement(nums1, nums2, totalLength/2+1))
	}
	mid1 := getKthElement(nums1, nums2, totalLength/2)
	mid2 := getKthElement(nums1, nums2, totalLength/2+1)
	return float64(mid1) + float64(mid2-mid1)/2
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
		half := k / 2
		newP1, newP2 := min(p1+half, len(nums1))-1, min(p2+half, len(nums2))-1
		if nums1[newP1] < nums2[newP2] {
			k -= newP1 - p1 + 1 // 这个+1是因为之后 p不是 newP1而是 newP1+1
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
	// 对较短数组没有使用二分，O(min(m,n)) 不合题意
	m, n := len(nums1), len(nums2)
	if m > n {
		return solution5(nums2, nums1)
	}
	totalLength := m + n
	if totalLength == 0 {
		return 0
	}
	hafl := (totalLength + 1) / 2
	// 设 nums1 的分界线左侧为 p1，nums2 的分界线左侧为 p2
	// 则有 (p1 + 1) + (p2 + 1) = half
	p1, p2 := m-1, hafl-m-1
	for p1 >= 0 && nums1[p1] >= nums2[p2+1] {
		p1--
		p2++
	}
	l1, r1, l2, r2 := math.MinInt, math.MaxInt, math.MinInt, math.MaxInt
	if p1 >= 0 {
		l1 = nums1[p1]
	}
	if p1+1 < m {
		r1 = nums1[p1+1]
	}
	if p2 >= 0 {
		l2 = nums2[p2]
	}
	if p2+1 < n {
		r2 = nums2[p2+1]
	}
	if totalLength&1 == 1 {
		return float64(max(l1, l2))
	}
	mid1, mid2 := max(l1, l2), min(r1, r2)
	return float64(mid1) + float64(mid2-mid1)/2
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
