package main

func numIslands(grid [][]byte) int {
	//return solution1(grid)
	return solution2(grid)
}

func solution1(grid [][]byte) int {
	// dfs
	if len(grid) < 1 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	var ans int
	for i := 0; i < m; i++ {
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

func dfs(grid [][]byte, m, n, x, y int) {
	if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == '0' {
		return
	}
	grid[x][y] = '0'
	dfs(grid, m, n, x-1, y)
	dfs(grid, m, n, x, y+1)
	dfs(grid, m, n, x+1, y)
	dfs(grid, m, n, x, y-1)
}

func solution2(grid [][]byte) int {
	// bfs
	if len(grid) < 1 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	var ans int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}
			bfs(grid, m, n, i, j)
			ans++
		}
	}
	return ans
}

func bfs(grid [][]byte, m, n, i, j int) {
	q := [][2]int{{i, j}}
	for len(q) > 0 {
		tmp := q[0]
		q = q[1:]
		x, y := tmp[0], tmp[1]
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == '0' {
			continue
		}
		grid[x][y] = '0'
		q = append(q, [2]int{x - 1, y})
		q = append(q, [2]int{x, y + 1})
		q = append(q, [2]int{x + 1, y})
		q = append(q, [2]int{x, y - 1})
	}
}
