package lc0703

import "math"

type KthLargest struct {
	heap []int
}

func Constructor(k int, nums []int) KthLargest {
	heap := make([]int, k)
	for i := range heap {
		heap[i] = math.MinInt
	}
	copy(heap, nums)
	buildMinHeap(heap)
	for i := k; i < len(nums); i++ {
		if nums[i] > heap[0] {
			heap[0] = nums[i]
			minHeapify(heap, 0)
		}
	}
	return KthLargest{heap: heap}
}

func (this *KthLargest) Add(val int) int {
	if val > this.heap[0] {
		this.heap[0] = val
		minHeapify(this.heap, 0)
	}
	return this.heap[0]
}

func buildMinHeap(heap []int) {
	for i := len(heap) >> 1; i >= 0; i-- {
		minHeapify(heap, i)
	}
}

func minHeapify(heap []int, i int) {
	l, r, m, n := i<<1, i<<1+1, i, len(heap)
	if l < n && heap[l] < heap[m] {
		m = l
	}
	if r < n && heap[r] < heap[m] {
		m = r
	}
	if m == i {
		return
	}
	heap[i], heap[m] = heap[m], heap[i]
	minHeapify(heap, m)
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
