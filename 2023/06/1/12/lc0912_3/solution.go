package lc0912

import (
	"math/rand"
	"time"
)

func sortArray(nums []int) []int {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	quickSort(rnd, nums, 0, len(nums)-1)
	return nums
}

func quickSort(rnd *rand.Rand, nums []int, l, r int) {
	if l >= r {
		return
	}
	p := partition(rnd, nums, l, r)
	quickSort(rnd, nums, l, p-1)
	quickSort(rnd, nums, p+1, r)
}

func partition(rnd *rand.Rand, nums []int, l, r int) int {
	tmp := l + rnd.Int()%(r-l+1)
	nums[l], nums[tmp] = nums[tmp], nums[l]
	pivot, lt, gt := nums[l], l+1, r
	for {
		for ; lt <= gt && nums[lt] < pivot; lt++ {
		}
		for ; lt <= gt && nums[gt] > pivot; gt-- {
		}
		if lt > gt {
			break
		}
		nums[lt], nums[gt] = nums[gt], nums[lt]
		lt++
		gt--
	}
	nums[l], nums[gt] = nums[gt], nums[l]
	return gt
}
