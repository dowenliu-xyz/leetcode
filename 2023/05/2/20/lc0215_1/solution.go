package lc0215_1

import (
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	//return solution1(nums, k)
	return solution2(nums, k)
}

func solution1(nums []int, k int) int {
	// min heap
	if len(nums) < k {
		return 0
	}
	heap := make([]int, k)
	copy(heap, nums)
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
	heap[i], heap[m] = heap[m], heap[i]
	minHeapify(heap, m, size)
}

func solution2(nums []int, k int) int {
	if len(nums) < k {
		return 0
	}
	target := len(nums) - k
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return quickSelect(rnd, nums, 0, len(nums)-1, target)
}

func quickSelect(rnd *rand.Rand, nums []int, left, right, target int) int {
	pivot := partition(rnd, nums, left, right)
	if pivot == target {
		return nums[pivot]
	}
	if pivot < target {
		return quickSelect(rnd, nums, pivot+1, right, target)
	}
	return quickSelect(rnd, nums, left, pivot-1, target)
}

func partition(rnd *rand.Rand, nums []int, left, right int) int {
	tmp := rnd.Int()%(right-left+1) + left
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
