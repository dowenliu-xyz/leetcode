package main

import (
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	//return solution1(nums, k)
	//return solution2(nums, k)
	return solution3(nums, k)
}

func solution1(nums []int, k int) int {
	// 建大顶堆再移除元素获取第K大元素
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := 0; i < k-1; i++ {
		heapSize--
		nums[0], nums[heapSize] = nums[heapSize], nums[0]
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

func buildMaxHeap(a []int, heapSize int) {
	for i := heapSize >> 1; i >= 0; i-- {
		maxHeapify(a, i, heapSize)
	}
}

func maxHeapify(a []int, i, heapSize int) {
	l, r, m := i<<1+1, i<<1+2, i
	if l < heapSize && a[l] > a[m] {
		m = l
	}
	if r < heapSize && a[r] > a[m] {
		m = r
	}
	if m != i {
		a[i], a[m] = a[m], a[i]
		maxHeapify(a, m, heapSize)
	}
}

func solution2(nums []int, k int) int {
	// 取前 K 元素建小顶堆，遍历剩下的元素，如果比堆顶大，替换堆顶并重新堆化
	heap := make([]int, k)
	copy(heap, nums[:k])
	buildMinHeap(heap, k)
	for i := k; i < len(nums); i++ {
		if nums[i] > heap[0] {
			heap[0] = nums[i]
			minHeapify(heap, 0, k)
		}
	}
	return heap[0]
}

func buildMinHeap(a []int, heapSize int) {
	for i := heapSize >> 1; i >= 0; i-- {
		minHeapify(a, i, heapSize)
	}
}

func minHeapify(a []int, i, heapSize int) {
	l, r, m := i<<1+1, i<<1+2, i
	if l < heapSize && a[l] < a[m] {
		m = l
	}
	if r < heapSize && a[r] < a[m] {
		m = r
	}
	if m != i {
		a[i], a[m] = a[m], a[i]
		minHeapify(a, m, heapSize)
	}
}

func solution3(nums []int, k int) int {
	// 快速选择算法
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(a []int, l, r, index int) int {
	p := randomPartition(a, l, r)
	if p == index {
		return a[p]
	}
	if p < index {
		return quickSelect(a, p+1, r, index)
	}
	return quickSelect(a, l, p-1, index)
}

func randomPartition(a []int, l, r int) int {
	i := rand.Int()%(r-l+1) + l
	a[i], a[r] = a[r], a[i]
	return partition(a, l, r)
}

func partition(a []int, l, r int) int {
	p := l - 1
	for i := l; i < r; i++ {
		if a[i] < a[r] {
			p++
			a[i], a[p] = a[p], a[i]
		}
	}
	a[p+1], a[r] = a[r], a[p+1]
	return p + 1
}
