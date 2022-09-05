package main

import "strings"

func readGrid(s string) [][]byte {
	var grid [][]byte
	s = strings.TrimSpace(s)
	s = strings.TrimLeft(s, "[")
	s = strings.TrimRight(s, "]")
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		line = strings.TrimLeft(line, "[")
		line = strings.TrimRight(line, "],")
		cs := strings.Split(line, ",")
		bs := make([]byte, 0, len(cs))
		for _, c := range cs {
			c = strings.TrimSpace(c)
			c = strings.Trim(c, "\"")
			bs = append(bs, []byte(c)[0])
		}
		grid = append(grid, bs)
	}
	return grid
}
