package lc0093_1

import "strings"

func restoreIpAddresses(s string) []string {
	var ans []string
	search([]byte(s), &ans, make([][]byte, 0, 4), 0) // 先把 s 转换为 []byte ，减少后续反复转换的成本
	return ans
}

func search(s []byte, ans *[]string, parts [][]byte, i int) {
	if len(parts) == 4 { // 找到一个结果
		*ans = append(*ans, join(parts))
		return
	}
	if len(s[i:]) < 4-len(parts) || len(s[i:]) > 3*(4-len(parts)) { // 剪枝：每段 1 - 3 位数字，不够分或分不开
		return
	}
	if len(parts) == 3 { // 最后一段要特别处理，不能枚举要 1 - 3 位，要把剩余的位数整体作为最后一段
		if isValidPart(s[i:]) { // 只有合法的段才加入
			search(s, ans, append(parts, s[i:]), 0)
		}
		return
	}
	for j := 1; j <= 3; j++ { // 其他段要枚举所有的可能，之后的 1 - 3 位是否能作为一段来用。
		if i+j > len(s) || !isValidPart(s[i:i+j]) { // 注意不能超出 s 的长度
			continue
		}
		search(s, ans, append(parts, s[i:i+j]), i+j)
	}
}

func isValidPart(s []byte) bool {
	n := len(s)
	if n > 1 && s[0] == '0' { // 除了 0 之外有前导 0，无效
		return false
	}
	// 手工转换数字，也可以使用 strconv.Atoi
	tmp := 0
	for i := 0; i < n; i++ {
		tmp = tmp*10 + int(s[i]-'0')
	}
	return tmp <= 255
}

func join(parts [][]byte) string {
	sb := strings.Builder{}
	for i, part := range parts {
		if i > 0 {
			sb.WriteByte('.')
		}
		sb.Write(part)
	}
	return sb.String()
}
