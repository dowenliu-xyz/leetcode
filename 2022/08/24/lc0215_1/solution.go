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

func solution1(a []int, k int) int {
	// 建立大顶堆，然后移除 k-1 次堆顶，之后堆顶即所得结果
	heapSize := len(a)
	buildMaxHeap(a, heapSize)
	for i := 0; i < k-1; i++ {
		a[0], a[heapSize-1] = a[heapSize-1], a[0]
		heapSize--
		maxHeapify(a, 0, heapSize)
	}
	return a[0]
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
	// 建立一个容量为 K 的小顶堆，将 nums 中比堆顶元素大的值 replace 入堆。最后堆顶即所求
	heap := make([]int, k)
	copy(heap, nums[:k])
	buildMinHeap(heap, k)
	for i := k; i < len(nums); i++ {
		if nums[i] <= heap[0] {
			continue
		}
		heap[0] = nums[i]
		minHeapify(heap, 0, k)
	}
	return heap[0]
}

func buildMinHeap(heap []int, heapSize int) {
	for i := heapSize >> 1; i >= 0; i-- {
		minHeapify(heap, i, heapSize)
	}
}

func minHeapify(heap []int, i, heapSize int) {
	l, r, m := i<<1+1, i<<1+2, i
	if l < heapSize && heap[l] < heap[m] {
		m = l
	}
	if r < heapSize && heap[r] < heap[m] {
		m = r
	}
	if m != i {
		heap[i], heap[m] = heap[m], heap[i]
		minHeapify(heap, m, heapSize)
	}
}

func solution3(nums []int, k int) int {
	// 快速选择算法
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
	x := nums[r]
	p := l - 1
	for i := l; i < r; i++ {
		if nums[i] < x {
			p++
			nums[p], nums[i] = nums[i], nums[p]
		}
	}
	nums[p+1], nums[r] = nums[r], nums[p+1]
	return p + 1
}
