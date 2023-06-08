package lc0069

func mySqrt(x int) int {
	low, high := 0, x
	for low <= high {
		mid := low + (high-low)>>1
		square := mid * mid
		if square == x {
			return mid
		}
		if square < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return high
}
