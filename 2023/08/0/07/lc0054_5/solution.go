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
	res := make([]int, 0, rows*cols)
	top, right, bottom, left := 0, cols-1, rows-1, 0
	for top <= bottom && left <= right {
		for col := left; col <= right; col++ {
			res = append(res, matrix[top][col])
		}
		for row := top + 1; row <= bottom; row++ {
			res = append(res, matrix[row][right])
		}
		if top < bottom {
			for col := right - 1; col >= left; col-- {
				res = append(res, matrix[bottom][col])
			}
		}
		if left < right {
			for row := bottom - 1; row > top; row-- {
				res = append(res, matrix[row][left])
			}
		}
		top++
		right--
		bottom--
		left++
	}
	return res
}
