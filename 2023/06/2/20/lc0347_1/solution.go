package lc0347

func topKFrequent(nums []int, k int) []int {
	// 最小堆解法。还可以用快速选择算法，时间复杂度 O(n)
	counts := map[int]int{}
	for _, num := range nums {
		counts[num]++
	}
	heap := make([]Stat, 0, k)
	for num, count := range counts {
		if len(heap) < k {
			heap = append(heap, Stat{Val: num, Count: count})
			continue
		}
		if len(heap) == k {
			buildMinHeap(heap)
		}
		if count > heap[0].Count {
			heap[0] = Stat{Val: num, Count: count}
			minHeapify(heap, 0)
		}
	}
	ans := make([]int, k)
	for i, s := range heap {
		ans[i] = s.Val
	}
	return ans
}

func buildMinHeap(heap []Stat) {
	for i := len(heap) >> 1; i >= 0; i-- {
		minHeapify(heap, i)
	}
}

func minHeapify(heap []Stat, i int) {
	l, r, m := i<<1+1, i<<1+2, i
	if l < len(heap) && heap[l].Count < heap[m].Count {
		m = l
	}
	if r < len(heap) && heap[r].Count < heap[m].Count {
		m = r
	}
	if m == i {
		return
	}
	heap[i], heap[m] = heap[m], heap[i]
	minHeapify(heap, m)
}

type Stat struct {
	Val   int
	Count int
}
