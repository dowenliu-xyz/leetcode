package lc0703_2

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
	buildMinHeap(heap, k)
	kl := KthLargest{heap: heap}
	for i := k; i < len(nums); i++ {
		kl.Add(nums[i])
	}
	return kl
}

func (this *KthLargest) Add(val int) int {
	if val > this.heap[0] {
		this.heap[0] = val
		minHeapify(this.heap, 0, len(this.heap))
	}
	return this.heap[0]
}

func buildMinHeap(heap []int, size int) {
	for i := size >> 1; i >= 0; i-- {
		minHeapify(heap, i, size)
	}
}

func minHeapify(heap []int, i, size int) {
	l, r, m := i<<1+1, i<<1+2, i
	if l < size && heap[l] < heap[m] {
		m = l
	}
	if r < size && heap[r] < heap[m] {
		m = r
	}
	if m == i {
		return
	}
	heap[m], heap[i] = heap[i], heap[m]
	minHeapify(heap, m, size)
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
