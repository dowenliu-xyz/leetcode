package main

import "sort"

func threeSum(nums []int) [][]int {
	//return solution1(nums)
	//return solution2(nums)
	//return solution3(nums)
	return solution4(nums)
}

func solution1(nums []int) [][]int {
	// 暴力枚举
	// O(n^3) 超时
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
					seen[a] = true
					ans = append(ans, a[:])
				}
			}
		}
	}
	return ans
}

func solution2(nums []int) [][]int {
	// 暴力枚举加查表优化。 O(n^2)
	var ans [][]int
	seen := map[[3]int]bool{}
	table := map[int]int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if k, ok := table[nums[i]+nums[j]]; ok {
				a := [3]int{nums[i], nums[j], nums[k]}
				sort.Ints(a[:])
				if seen[a] {
					continue
				}
				seen[a] = true
				ans = append(ans, a[:])
			}
		}
		table[-nums[i]] = i
	}
	return ans
}

func solution3(nums []int) [][]int {
	// 排序+双边收敛
	sort.Ints(nums)
	var ans [][]int
	seen := map[[3]int]bool{}
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				a := [3]int{nums[i], nums[left], nums[right]}
				if !seen[a] {
					seen[a] = true
					ans = append(ans, a[:])
				}
				left++
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return ans
}

func solution4(nums []int) [][]int {
	// 排序+双边收敛 优化
	sort.Ints(nums)
	var ans [][]int
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				for ; left < right && nums[left] == nums[left+1]; left++ {
				}
				for ; left < right && nums[right] == nums[right-1]; right-- {
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return ans
}
