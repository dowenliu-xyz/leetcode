package lc0006_2

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	rows := make([][]byte, numRows)
	row, dir := 0, -1
	n := len(s)
	for i := 0; i < n; i++ {
		rows[row] = append(rows[row], s[i])
		if row == 0 || row == numRows-1 {
			dir = -dir
		}
		row += dir
	}
	ans := make([]byte, 0, n)
	for i := range rows {
		ans = append(ans, rows[i]...)
	}
	return string(ans)
}
