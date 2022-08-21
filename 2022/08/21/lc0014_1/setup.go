package main

import "strings"

func parseStrings(input string) []string {
	input = strings.TrimSpace(input)
	input = strings.TrimLeft(input, "[")
	input = strings.TrimRight(input, "]")
	parts := strings.Split(input, ",")
	res := make([]string, 0, len(parts))
	for _, s := range parts {
		s = strings.TrimSpace(s)
		s = strings.Trim(s, `"`)
		res = append(res, s)
	}
	return res
}
