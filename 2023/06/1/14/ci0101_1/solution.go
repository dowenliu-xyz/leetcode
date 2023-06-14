package ci0101

func isUnique(astr string) bool {
	var used uint32
	for i := 0; i < len(astr); i++ {
		p := astr[i] - 'a'
		if used&(1<<p) != 0 {
			return false
		}
		used |= 1 << p
	}
	return true
}
