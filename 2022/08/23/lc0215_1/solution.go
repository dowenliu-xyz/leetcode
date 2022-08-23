package main

import (
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, l, r, index int) int {
	q := randomPartition(nums, l, r)
	if q == index {
		return nums[q]
	}
	if q < index {
		return quickSelect(nums, q+1, r, index)
	}
	return quickSelect(nums, l, q-1, index)
}

func randomPartition(nums []int, l, r int) int {
	i := rand.Int()%(r-l+1) + l
	nums[i], nums[r] = nums[r], nums[i]
	return partition(nums, l, r)
}

func partition(nums []int, l, r int) int {
	p := l - 1
	for i := l; i < r; i++ {
		if nums[i] < nums[r] {
			p++
			nums[p], nums[i] = nums[i], nums[p]
		}
	}
	nums[p+1], nums[r] = nums[r], nums[p+1]
	return p + 1
}
