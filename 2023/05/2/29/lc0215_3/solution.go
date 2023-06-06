package lc0215_3

import (
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	return solution1(nums, k)
	//return solution2(nums, k)
}

func solution1(nums []int, k int) int {
	n := len(nums)
	if n < k {
		return -1
	}
	heap := make([]int, k)
	copy(heap, nums)
	buildMinHeap(heap, k)
	for i := k; i < n; i++ {
		if nums[i] > heap[0] {
			heap[0] = nums[i]
			minHeapify(heap, 0, k)
		}
	}
	return heap[0]
}

func buildMinHeap(heap []int, size int) {
	for i := size >> 1; i >= 0; i-- {
		minHeapify(heap, i, size)
	}
}

func minHeapify(heap []int, i, size int) {
	l, r, m := i<<1, i<<1+1, i
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

func solution2(nums []int, k int) int {
	if len(nums) < k {
		return -1
	}
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return quickSortSelect(rnd, nums, 0, len(nums)-1, len(nums)-k)
}

func quickSortSelect(rnd *rand.Rand, nums []int, left, right, target int) int {
	pivot := randomPartition(rnd, nums, left, right)
	if pivot == target {
		return nums[pivot]
	}
	if pivot < target {
		return quickSortSelect(rnd, nums, pivot+1, right, target)
	}
	return quickSortSelect(rnd, nums, left, pivot-1, target)
}

func randomPartition(rnd *rand.Rand, nums []int, left, right int) int {
	tmp := left + rnd.Int()%(right-left+1)
	nums[tmp], nums[right] = nums[right], nums[tmp]
	p := left - 1
	for i := left; i < right; i++ {
		if nums[i] < nums[right] {
			p++
			nums[i], nums[p] = nums[p], nums[i]
		}
	}
	p++
	nums[p], nums[right] = nums[right], nums[p]
	return p
}
