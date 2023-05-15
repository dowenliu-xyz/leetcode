package lc0054_2

func spiralOrder(matrix [][]int) []int {
	//return solution1(matrix)
	return solution2(matrix)
}

func solution1(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	if n == 0 {
		return nil
	}
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	target := m * n
	d, x, y, ans := 0, 0, 0, make([]int, 0, target)
	for i := 0; i < target; i++ {
		ans = append(ans, matrix[x][y])
		vis[x][y] = true
		newX, newY := x+dir[d][0], y+dir[d][1]
		if newX < 0 || newX >= m || newY < 0 || newY >= n || vis[newX][newY] {
			d = (d + 1) & 3
		}
		x, y = x+dir[d][0], y+dir[d][1]
	}
	return ans
}

func solution2(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	if n == 0 {
		return nil
	}
	ans := make([]int, 0, m*n)
	top, right, bottom, left := 0, n-1, m-1, 0
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
