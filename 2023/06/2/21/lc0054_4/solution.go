package lc0054

func spiralOrder(matrix [][]int) []int {
	rows := len(matrix)
	if rows == 0 {
		return nil
	}
	cols := len(matrix[0])
	if cols == 0 {
		return nil
	}
	top, right, bottom, left, ans := 0, cols-1, rows-1, 0, make([]int, 0, rows*cols)
	for top <= bottom && left <= right {
		for col := left; col <= right; col++ {
			ans = append(ans, matrix[top][col])
		}
		for row := top + 1; row <= bottom; row++ {
			ans = append(ans, matrix[row][right])
		}
		if top < bottom {
			for col := right - 1; col >= left; col-- {
				ans = append(ans, matrix[bottom][col])
			}
		}
		if left < right {
			for row := bottom - 1; row > top; row-- {
				ans = append(ans, matrix[row][left])
			}
		}
		top++
		right--
		bottom--
		left++
	}
	return ans
}
