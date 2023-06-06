package lc0054_3

func spiralOrder(matrix [][]int) []int {
	//return solution1(matrix)
	return solution2(matrix)
}

func solution1(matrix [][]int) []int {
	n := len(matrix)
	if n == 0 {
		return nil
	}
	m := len(matrix[0])
	if m == 0 {
		return nil
	}
	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, m)
	}
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	x, y, dir, total := 0, 0, 0, m*n
	ans := make([]int, total)
	for i := 0; i < total; i++ {
		ans[i] = matrix[x][y]
		vis[x][y] = true
		newX, newY := x+directions[dir][0], y+directions[dir][1]
		if newX < 0 || newX >= n || newY < 0 || newY >= m || vis[newX][newY] {
			dir = (dir + 1) & 3
		}
		x, y = x+directions[dir][0], y+directions[dir][1]
	}
	return ans
}

func solution2(matrix [][]int) []int {
	rows := len(matrix)
	if rows == 0 {
		return nil
	}
	cols := len(matrix[0])
	if cols == 0 {
		return nil
	}
	top, right, bottom, left := 0, cols-1, rows-1, 0
	ans := make([]int, 0, rows*cols)
	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[top][i])
		}
		for i := top + 1; i <= bottom; i++ {
			ans = append(ans, matrix[i][right])
		}
		if top < bottom {
			for i := right - 1; i >= left; i-- {
				ans = append(ans, matrix[bottom][i])
			}
		}
		if left < right {
			for i := bottom - 1; i > top; i-- {
				ans = append(ans, matrix[i][left])
			}
		}
		top++
		right--
		bottom--
		left++
	}
	return ans
}
