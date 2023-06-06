package lc0064_2

func minPathSum(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	if cols == 0 {
		return 0
	}
	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < cols; i++ {
		dp[0][i] = grid[0][i] + dp[0][i-1]
	}
	for i := 1; i < rows; i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
		for j := 1; j < cols; j++ {
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[rows-1][cols-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
