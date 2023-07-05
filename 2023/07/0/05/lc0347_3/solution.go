package lc0347

import (
	"math/rand"
	"time"
)

func topKFrequent(nums []int, k int) []int {
	//return solution1(nums, k)
	return solution2(nums, k)
}

func solution1(nums []int, k int) []int {
	stats := frequentStats(nums)
	heap := make([][]int, k)
	copy(heap, stats)
	var minHeapify func(heap [][]int, i int)
	minHeapify = func(heap [][]int, i int) {
		l, r, m, s := i<<1+1, i<<1+2, i, len(heap)
		if l < s && heap[l][1] < heap[m][1] {
			m = l
		}
		if r < s && heap[r][1] < heap[m][1] {
			m = r
		}
		if m == i {
			return
		}
		heap[i], heap[m] = heap[m], heap[i]
		minHeapify(heap, m)
	}
	buildMinHeap := func(heap [][]int) {
		for i := len(heap) >> 1; i >= 0; i-- {
			minHeapify(heap, i)
		}
	}
	buildMinHeap(heap)
	for i := k; i < len(stats); i++ {
		if stats[i][1] > heap[0][1] {
			heap[0] = stats[i]
			minHeapify(heap, 0)
		}
	}
	ans := make([]int, 0, k)
	for i := range heap {
		ans = append(ans, heap[i][0])
	}
	return ans
}

func solution2(nums []int, k int) []int {
	stats := frequentStats(nums)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	partition := func(left, right int) int {
		tmp := left + rnd.Int()%(right+1-left)
		stats[left], stats[tmp] = stats[tmp], stats[left]
		l, r, p := left+1, right, stats[left][1]
		for l <= r {
			for ; l <= r && stats[l][1] > p; l++ {
			}
			for ; l <= r && stats[r][1] <= p; r-- {
			}
			if l > r {
				break
			}
			stats[l], stats[r] = stats[r], stats[l]
			l++
			r--
		}
		stats[left], stats[r] = stats[r], stats[left]
		return r
	}
	left, right := 0, len(stats)-1
	for left <= right {
		p := partition(left, right)
		if p == k-1 {
			ans := make([]int, 0, k)
			for i := range stats[:k] {
				ans = append(ans, stats[i][0])
			}
			return ans
		}
		if p < k-1 {
			left = p + 1
		} else {
			right = p - 1
		}
	}
	return nil
}

func frequentStats(nums []int) [][]int {
	mp := map[int]int{}
	for _, num := range nums {
		mp[num]++
	}
	stats := make([][]int, 0, len(mp))
	for num, count := range mp {
		stats = append(stats, []int{num, count})
	}
	return stats
}
