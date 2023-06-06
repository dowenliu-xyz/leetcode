package lc0051_3

import "math/bits"

func solveNQueens(n int) [][]string {
	if n == 0 {
		return nil
	}
	var ans [][]string
	solve(&ans, n, make([]int, 0, n), 0, 0, 0)
	return ans
}

func solve(ans *[][]string, n int, queens []int, colBits, pieBits, naBits uint) {
	if len(queens) == n {
		*ans = append(*ans, generateBoard(queens, n))
		return
	}
	invalidBits := colBits | pieBits | naBits
	possibleBits := (uint(1)<<n - 1) & (^invalidBits)
	for possibleBits != 0 {
		pos := possibleBits & (-possibleBits)
		possibleBits = possibleBits & (possibleBits - 1)
		col := bits.OnesCount(pos - 1)
		solve(ans, n, append(queens, col), colBits|pos, (pieBits|pos)<<1, (naBits|pos)>>1)
	}
}

func generateBoard(queens []int, n int) []string {
	board := make([]string, n)
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board[i] = string(row)
	}
	return board
}
