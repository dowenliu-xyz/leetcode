package lc0051_1

import "math/bits"

func solveNQueens(n int) [][]string {
	var ans [][]string
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	solve(&ans, queens, n, 0, 0, 0, 0)
	return ans
}

func solve(ans *[][]string, queens []int, n, row int, columns, diagonal1, diagonal2 uint32) {
	if n == row {
		*ans = append(*ans, generateBoard(queens, n))
		return
	}
	unavailablePositions := columns | diagonal1 | diagonal2    // 所有不能放的位置
	allPositions := (uint32(1) << n) - 1                       // 行内所有位置
	availablePositions := allPositions & ^unavailablePositions // 所有可以放的位置
	for availablePositions != 0 {
		position := availablePositions & (-availablePositions)             // 取最低位 1
		availablePositions = availablePositions & (availablePositions - 1) // 清除最低位的 1
		column := bits.OnesCount32(position - 1)                           // 当前位置列号。position - 1，当前位置低位全1，count 后正好是位置列号
		queens[row] = column
		solve(ans, queens, n, row+1, columns|position, (diagonal1|position)<<1, (diagonal2|position)>>1)
		queens[row] = -1
	}
}

func generateBoard(queens []int, n int) []string {
	board := make([]string, 0, len(queens))
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}
