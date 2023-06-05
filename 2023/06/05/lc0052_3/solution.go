package lc0052

import "math/bits"

func totalNQueens(n int) int {
	if n < 1 {
		return 0
	}
	ans := 0
	var resolve func(queens []int, colBits, pieBits, naBits uint)
	resolve = func(queens []int, colBits, pieBits, naBits uint) {
		if len(queens) == n {
			ans++
			return
		}
		invalidBits := colBits | pieBits | naBits
		availableBits := (uint(1)<<n - 1) & (^invalidBits)
		for availableBits != 0 {
			pos := availableBits & (-availableBits)
			availableBits = availableBits & (availableBits - 1)
			col := bits.OnesCount(pos - 1)
			resolve(append(queens, col), colBits|pos, (pieBits|pos)<<1, (naBits|pos)>>1)
		}
	}
	resolve(make([]int, 0, n), 0, 0, 0)
	return ans
}
