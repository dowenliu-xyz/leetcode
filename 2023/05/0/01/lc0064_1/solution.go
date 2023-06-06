package lc0064_1

func minPathSum(grid [][]int) int {
	m := len(grid)
	if m < 1 {
		return 0
	}
	n := len(grid[0])
	if n < 1 {
		return 0
	}
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}
	f[0][0] = grid[0][0]
	for i := 1; i < n; i++ {
		f[0][i] = f[0][i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		f[i][0] = f[i-1][0] + grid[i][0]
		for j := 1; j < n; j++ {
			f[i][j] = min(f[i-1][j], f[i][j-1]) + grid[i][j]
		}
	}
	return f[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
