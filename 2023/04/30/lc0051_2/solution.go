package lc0051_2

import "math/bits"

func solveNQueens(n int) [][]string {
	var ans [][]string
	queens := make([]int, n)
	for i := range queens {
		queens[i] = -1
	}
	solve(&ans, queens, n, 0, 0, 0, 0)
	return ans
}

func solve(ans *[][]string, queens []int, n, row int, cBits, tlbBits, trbBits uint32) {
	if n == row {
		*ans = append(*ans, generateBoard(queens, n))
		return
	}
	invalidPositionBits := cBits | tlbBits | trbBits
	allPositionBits := (uint32(1) << n) - 1
	availablePositionBits := allPositionBits & (^invalidPositionBits)
	for availablePositionBits != 0 {
		position := availablePositionBits & (-availablePositionBits)
		availablePositionBits = availablePositionBits & (availablePositionBits - 1)
		column := bits.OnesCount32(position - 1)
		queens[row] = column
		solve(ans, queens, n, row+1, cBits|position, (tlbBits|position)<<1, (trbBits|position)>>1)
		queens[row] = -1
	}
}

func generateBoard(queens []int, n int) []string {
	board := make([]string, 0, n)
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
