package lc0278_2

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	return binarySearch(1, n)
}

func binarySearch(l, r int) int {
	m := l + (r-l)>>1
	if isBadVersion(m) {
		if m == 1 || !isBadVersion(m-1) {
			return m
		}
		return binarySearch(l, m-1)
	}
	return binarySearch(m+1, r)
}
