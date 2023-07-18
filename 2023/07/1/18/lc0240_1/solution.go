package lc0240

func searchMatrix(matrix [][]int, target int) bool {
	//return solution1(matrix, target)
	return solution2(matrix, target)
}

func solution1(matrix [][]int, target int) bool {
	rows := len(matrix)
	if rows == 0 {
		return false
	}
	cols := len(matrix[0])
	if cols == 0 {
		return false
	}
	binarySearch := func(nums []int, target int) bool {
		left, right := 0, len(nums)-1
		for left <= right {
			m := left + (right-left)>>1
			if nums[m] == target {
				return true
			}
			if nums[m] < target {
				left = m + 1
			} else {
				right = m - 1
			}
		}
		return false
	}
	for _, nums := range matrix {
		if binarySearch(nums, target) {
			return true
		}
	}
	return false
}

func solution2(matrix [][]int, target int) bool {
	rows := len(matrix)
	if rows == 0 {
		return false
	}
	cols := len(matrix[0])
	if cols == 0 {
		return false
	}
	row, col := 0, cols-1
	for row < rows && col >= 0 {
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] < target {
			row++
		} else {
			col--
		}
	}
	return false
}
