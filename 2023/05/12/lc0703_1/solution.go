package lc0703_1

import "math"

type KthLargest struct {
	heap []int
}

func Constructor(k int, nums []int) KthLargest {
	heap := make([]int, k+1)
	for i := range heap {
		heap[i] = math.MinInt
	}
	copy(heap[1:], nums)
	buildMinHeap(heap)
	l := KthLargest{
		heap: heap,
	}
	for i := k; i < len(nums); i++ {
		l.Add(nums[i])
	}
	return l
}

func (this *KthLargest) Add(val int) int {
	if val > this.heap[1] {
		this.heap[1] = val
		minHeapify(this.heap, 1)
	}
	return this.heap[1]
}

func buildMinHeap(heap []int) {
	for i := len(heap) >> 1; i > 0; i-- {
		minHeapify(heap, i)
	}
}

func minHeapify(heap []int, i int) {
	l, r, m := i<<1, i<<1+1, i
	if l < len(heap) && heap[l] < heap[m] {
		m = l
	}
	if r < len(heap) && heap[r] < heap[m] {
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
