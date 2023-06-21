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
	stat := map[int]int{}
	for _, num := range nums {
		stat[num]++
	}
	type Freq struct {
		Num   int
		Count int
	}
	var minHeapify func(heap []Freq, i int)
	minHeapify = func(heap []Freq, i int) {
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
	buildMinHeap := func(heap []Freq) {
		for i := len(heap) >> 1; i >= 0; i-- {
			minHeapify(heap, i)
		}
	}
	heap := make([]Freq, 0, k)
	for num, count := range stat {
		if len(heap) < k {
			heap = append(heap, Freq{Num: num, Count: count})
			continue
		}
		if len(heap) == k {
			buildMinHeap(heap)
		}
		if count > heap[0].Count {
			heap[0] = Freq{Num: num, Count: count}
			minHeapify(heap, 0)
		}
	}
	ans := make([]int, 0, k)
	for i := range heap {
		ans = append(ans, heap[i].Num)
	}
	return ans
}

func solution2(nums []int, k int) []int {
	statMap := map[int]int{}
	for _, num := range nums {
		statMap[num]++
	}
	stats := make([][]int, 0, len(statMap))
	for num, count := range statMap {
		stats = append(stats, []int{num, count})
	}
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	var partition func(stats [][]int, left, right int) int
	partition = func(stats [][]int, left, right int) int {
		tmp := left + rnd.Int()%(right-left+1)
		stats[left], stats[tmp] = stats[tmp], stats[left]
		gt, lt, pivot := left+1, right, stats[left][1]
		for gt <= lt {
			for ; gt <= lt && stats[gt][1] > pivot; gt++ {
			}
			for ; gt <= lt && stats[lt][1] <= pivot; lt-- {
			}
			if gt > lt {
				break
			}
			stats[gt], stats[lt] = stats[lt], stats[gt]
			gt++
			lt--
		}
		stats[left], stats[lt] = stats[lt], stats[left]
		return lt
	}
	left, right := 0, len(stats)-1
	for left <= right {
		p := partition(stats, left, right)
		if p == k-1 {
			ans := make([]int, 0, k)
			for _, stat := range stats[:k] {
				ans = append(ans, stat[0])
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
