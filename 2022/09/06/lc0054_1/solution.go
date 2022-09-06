package main

func spiralOrder(matrix [][]int) []int {
	//return solution1(matrix)
	return solution2(matrix)
}

func solution1(matrix [][]int) []int {
	// 记录访问位置和方向
	m := len(matrix)
	if m < 1 {
		return nil
	}
	n := len(matrix[0])
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	total := m * n
	ans := make([]int, total)
	d, x, y := 0, 0, 0
	for i := 0; i < total; i++ {
		ans[i] = matrix[x][y]
		visited[x][y] = true
		newX, newY := x+directions[d][0], y+directions[d][1]
		if newX < 0 || newX >= m || newY < 0 || newY >= n || visited[newX][newY] {
			d = (d + 1) % 4
		}
		x += directions[d][0]
		y += directions[d][1]
	}
	return ans
}

func solution2(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	m, n := len(matrix), len(matrix[0])
	left, right, top, bottom := 0, n-1, 0, m-1
	ans := make([]int, 0, m*n)
	for left <= right && top <= bottom {
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
		left++
		right--
		top++
		bottom--
	}
	return ans
}
