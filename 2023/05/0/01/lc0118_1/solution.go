package lc0118_1

func generate(numRows int) [][]int {
	if numRows < 1 {
		return nil
	}
	ans := make([][]int, numRows)
	ans[0] = []int{1}
	for i := 1; i < numRows; i++ {
		line := make([]int, i+1)
		line[0], line[i] = 1, 1
		for j := 1; j < i; j++ {
			line[j] = ans[i-1][j-1] + ans[i-1][j]
		}
		ans[i] = line
	}
	return ans
}
