package lc0052_1

import "math/bits"

func totalNQueens(n int) int {
	if n == 0 {
		return 0
	}
	return solve(n, make([]int, 0, n), 0, 0, 0)
}

func solve(n int, queens []int, colBits, pieBits, naBits uint) int {
	if n == len(queens) {
		return 1
	}
	ans := 0
	invalidBits := colBits | pieBits | naBits
	possibleBits := (uint(1)<<n - 1) & (^invalidBits)
	for possibleBits != 0 {
		pos := possibleBits & (-possibleBits)
		possibleBits = possibleBits & (possibleBits - 1)
		col := bits.OnesCount(pos - 1)
		ans += solve(n, append(queens, col), colBits|pos, (pieBits|pos)<<1, (naBits|pos)>>1)
	}
	return ans
}
