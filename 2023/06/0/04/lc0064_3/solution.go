package lc0064_3

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
	for col := 1; col < cols; col++ {
		dp[0][col] = dp[0][col-1] + grid[0][col]
	}
	for row := 1; row < rows; row++ {
		dp[row][0] = dp[row-1][0] + grid[row][0]
		for col := 1; col < cols; col++ {
			dp[row][col] = min(dp[row-1][col], dp[row][col-1]) + grid[row][col]
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
