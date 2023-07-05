package lc0703

type KthLargest struct {
	heap []int
	cap  int
}

func Constructor(k int, nums []int) KthLargest {
	heap := make([]int, 0, k)
	for _, num := range nums {
		if len(heap) < k {
			heap = appendToMinHeap(heap, num)
		} else if num > heap[0] {
			heap[0] = num
			minHeapify(heap, 0)
		}
	}
	return KthLargest{heap: heap, cap: k}
}

func (this *KthLargest) Add(val int) int {
	if len(this.heap) < this.cap {
		this.heap = appendToMinHeap(this.heap, val)
	} else if val > this.heap[0] {
		this.heap[0] = val
		minHeapify(this.heap, 0)
	}
	return this.heap[0]
}

func appendToMinHeap(heap []int, val int) []int {
	heap = append(heap, val)
	i := len(heap) - 1
	for i > 0 {
		p := (i - 1) >> 1
		if heap[p] <= heap[i] {
			break
		}
		heap[i], heap[p] = heap[p], heap[i]
		i = p
	}
	return heap
}

func minHeapify(heap []int, i int) {
	l, r, m, s := i<<1+1, i<<1+2, i, len(heap)
	if l < s && heap[l] < heap[m] {
		m = l
	}
	if r < s && heap[r] < heap[m] {
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
