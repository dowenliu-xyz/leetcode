package lc0052

func totalNQueens(n int) int {
	if n <= 1 {
		return n
	}
	ans := 0
	var dfs func(row int, colBits, pieBits, naBits uint)
	dfs = func(row int, colBits, pieBits, naBits uint) {
		if row == n {
			ans++
			return
		}
		availablePosBits := (uint(1)<<n - 1) & (^(colBits | pieBits | naBits))
		for i := 0; i < n; i++ {
			posBit := uint(1) << i
			if availablePosBits&posBit == 0 {
				continue
			}
			dfs(row+1, colBits|posBit, (pieBits|posBit)<<1, (naBits|posBit)>>1)
		}
	}
	dfs(0, 0, 0, 0)
	return ans
}
