package lc0059

func generateMatrix(n int) [][]int {
	if n == 0 {
		return nil
	}
	top, bottom, left, right := 0, n-1, 0, n-1
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	i := 1
	for top <= bottom && left <= right {
		for col := left; col <= right; col++ {
			ans[top][col] = i
			i++
		}
		for row := top + 1; row <= bottom; row++ {
			ans[row][right] = i
			i++
		}
		if bottom > top {
			for col := right - 1; col >= left; col-- {
				ans[bottom][col] = i
				i++
			}
			for row := bottom - 1; row > top; row-- {
				ans[row][left] = i
				i++
			}
		}
		top++
		right--
		bottom--
		left++
	}
	return ans
}
