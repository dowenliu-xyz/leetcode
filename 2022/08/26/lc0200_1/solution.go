package main

func numIslands(grid [][]byte) int {
	//return solution1(grid)
	return solution2(grid)
}

func solution1(grid [][]byte) int {
	// dfs 原地操作。如果要求不允许编辑输入数组就复制一遍，在副本上操作
	m := len(grid)
	if m == 0 {
		return 0
	}
	var ans int
	for i := 0; i < m; i++ {
		n := len(grid[i])
		if n == 0 {
			return 0
		}
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}
			dfs(grid, m, n, i, j)
			ans++
		}
	}
	return ans
}

func dfs(grid [][]byte, m, n, i, j int) {
	if grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	if i > 0 {
		dfs(grid, m, n, i-1, j)
	}
	if j+1 < n {
		dfs(grid, m, n, i, j+1)
	}
	if i+1 < m {
		dfs(grid, m, n, i+1, j)
	}
	if j > 0 {
		dfs(grid, m, n, i, j-1)
	}
}

func solution2(grid [][]byte) int {
	// bfs
	var ans int
	m := len(grid)
	if m == 0 {
		return 0
	}
	for i := 0; i < m; i++ {
		n := len(grid[i])
		if n == 0 {
			return 0
		}
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				bfs(grid, m, n, i, j)
				ans++
			}
		}
	}
	return ans
}

func bfs(grid [][]byte, m, n, i, j int) {
	q := [][2]int{{i, j}}
	for len(q) > 0 {
		c := q[0]
		q = q[1:]
		x, y := c[0], c[1]
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == '0' {
			continue
		}
		grid[x][y] = '0'
		q = append(q, [2]int{x - 1, y}, [2]int{x, y + 1}, [2]int{x + 1, y}, [2]int{x, y - 1})
	}
}
