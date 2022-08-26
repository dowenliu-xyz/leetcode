package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestLRUCache(t *testing.T) {
	tests := []struct {
		ops  string
		args string
		want string
	}{
		{
			ops:  `["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]`,
			args: `[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]`,
			want: `[null, null, null, 1, null, -1, null, -1, 3, 4]`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ops := readStrings(tt.ops, ",")
			args := parseArgs(readStrings(tt.args, "], ["))
			var results []string
			var obj *LRUCache
			for k, op := range ops {
				var result string
				obj, result = opMap[op](obj, args[k])
				results = append(results, result)
			}
			got := formatResults(results)
			want := strings.ReplaceAll(tt.want, " ", "")
			if want != got {
				t.Errorf("want %s, got %s", want, got)
			}
		})
	}
}
