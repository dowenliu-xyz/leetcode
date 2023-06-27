package lc0051

import "math/bits"

func solveNQueens(n int) [][]string {
	var ans [][]string
	var solve func(queens []int, colBits, pieBits, naBits uint)
	solve = func(queens []int, colBits, pieBits, naBits uint) {
		if len(queens) == n {
			ans = append(ans, generateBoard(queens, n))
			return
		}
		availableBits := (uint(1)<<n - 1) & (^(colBits | pieBits | naBits))
		for availableBits != 0 {
			pos := availableBits & (-availableBits)
			availableBits = availableBits & (availableBits - 1)
			col := bits.OnesCount(pos - 1)
			solve(append(queens, col), colBits|pos, (pieBits|pos)<<1, (naBits|pos)>>1)
		}
	}
	solve(make([]int, 0, n), 0, 0, 0)
	return ans
}

func generateBoard(queens []int, n int) []string {
	board := make([]string, 0, n)
	for _, queen := range queens {
		row := make([]byte, n)
		for i := range row {
			row[i] = '.'
		}
		row[queen] = 'Q'
		board = append(board, string(row))
	}
	return board
}
