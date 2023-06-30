package lc0215

import (
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	var partition func(left, right int) int
	partition = func(left, right int) int {
		if left >= right {
			return left
		}
		tmp := left + rnd.Int()%(right-left+1)
		nums[left], nums[tmp] = nums[tmp], nums[left]
		pivot, l, r := nums[left], left+1, right
		for l <= r {
			for ; l <= r && nums[l] < pivot; l++ {
			}
			for ; l <= r && pivot <= nums[r]; r-- {
			}
			if l > r {
				break
			}
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
		nums[left], nums[r] = nums[r], nums[left]
		return r
	}
	left, right, target := 0, len(nums)-1, len(nums)-k
	for {
		p := partition(left, right)
		if p == target {
			return nums[p]
		}
		if p < target {
			left = p + 1
		} else {
			right = p - 1
		}
	}
}
