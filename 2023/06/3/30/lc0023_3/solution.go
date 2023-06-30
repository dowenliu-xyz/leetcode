package lc0023

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	heap := make([]*ListNode, 0, len(lists))
	for _, list := range lists {
		if list == nil {
			continue
		}
		heap = append(heap, list)
	}
	buildMinHeap(heap)
	dummy := &ListNode{}
	pre := dummy
	for len(heap) > 0 {
		pre.Next, pre, heap[0] = heap[0], heap[0], heap[0].Next
		if heap[0] == nil {
			heap[0] = heap[len(heap)-1]
			heap = heap[:len(heap)-1]
		}
		minHeapify(heap, 0)
	}
	return dummy.Next
}

func buildMinHeap(heap []*ListNode) {
	for i := len(heap) >> 1; i >= 0; i-- {
		minHeapify(heap, i)
	}
}

func minHeapify(heap []*ListNode, i int) {
	l, r, m, size := i<<1+1, i<<1+2, i, len(heap)
	if l < size && heap[l].Val < heap[m].Val {
		m = l
	}
	if r < size && heap[r].Val < heap[m].Val {
		m = r
	}
	if m == i {
		return
	}
	heap[i], heap[m] = heap[m], heap[i]
	minHeapify(heap, m)
}
