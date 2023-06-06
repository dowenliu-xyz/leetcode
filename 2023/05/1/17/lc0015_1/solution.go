package lc0015_1

import "sort"

func threeSum(nums []int) [][]int {
	//return solution1(nums)
	return solution2(nums)
}

func solution1(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}
	var ans [][]int
	added := map[[3]int]bool{}
	mem := map[int]bool{nums[0]: true}
	for i := 1; i < n; i++ {
		for j := i + 1; j < n; j++ {
			k := 0 - nums[i] - nums[j]
			if mem[k] {
				array := [3]int{k, nums[i], nums[j]}
				sort.Ints(array[:])
				if added[array] {
					continue
				}
				ans = append(ans, array[:])
				added[array] = true
			}
		}
		mem[nums[i]] = true
	}
	return ans
}

func solution2(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}
	var ans [][]int
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				for ; l < r && nums[l] == nums[l+1]; l++ {
				}
				for ; l < r && nums[r-1] == nums[r]; r-- {
				}
				l++
				r--
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return ans
}
