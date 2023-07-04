package ci0101

func isUnique(astr string) bool {
	n := len(astr)
	if n == 0 {
		return true
	}
	if n > 26 {
		return false
	}
	seen := uint(0)
	for i := 0; i < n; i++ {
		bit := uint(1) << (astr[i] - 'a')
		if seen&bit != 0 {
			return false
		}
		seen |= bit
	}
	return true
}
