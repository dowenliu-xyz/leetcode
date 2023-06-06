package lc0006_1

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	rows := make([][]byte, numRows)
	row, flag := 0, -1
	for i := 0; i < len(s); i++ {
		rows[row] = append(rows[row], s[i])
		if row == 0 || row == numRows-1 {
			flag = -flag
		}
		row += flag
	}
	ans := make([]byte, 0, len(s))
	for i := range rows {
		ans = append(ans, rows[i]...)
	}
	return string(ans)
}
