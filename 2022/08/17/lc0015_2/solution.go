package main

import "sort"

func threeSum(nums []int) [][]int {
	//return solution1(nums)
	//return solution2(nums)
	return solution3(nums)
}

func solution1(nums []int) [][]int {
	// 暴力枚举 O(n^3) 超时
	var ans [][]int
	seen := map[[3]int]bool{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					a := [3]int{nums[i], nums[j], nums[k]}
					sort.Ints(a[:])
					if seen[a] {
						continue
					}
					ans = append(ans, a[:])
					seen[a] = true
				}
			}
		}
	}
	return ans
}

func solution2(nums []int) [][]int {
	// 暴力，+查表优化 O(n^2)
	var ans [][]int
	seen := map[[3]int]bool{}
	table := map[int]int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if k, ok := table[nums[i]+nums[j]]; ok {
				a := [3]int{nums[i], nums[j], nums[k]}
				sort.Ints(a[:])
				if !seen[a] {
					ans = append(ans, a[:])
					seen[a] = true
				}
			}
		}
		table[-nums[i]] = i
	}
	return ans
}

func solution3(nums []int) [][]int {
	// 排序后双边收敛
	var ans [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				for ; l < r && nums[l] == nums[l+1]; l++ {
				}
				for ; l < r && nums[r] == nums[r-1]; r-- {
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
