package main

import (
	"fmt"
	"strconv"
	"strings"
)

func readStrings(input, sep string) []string {
	input = strings.TrimSpace(input)
	input = strings.TrimLeft(input, "[")
	input = strings.TrimRight(input, "]")
	parts := strings.Split(input, sep)
	res := make([]string, 0, len(parts))
	for _, s := range parts {
		trim := strings.TrimSpace(s)
		if "" != trim {
			res = append(res, trim)
		}
	}
	return res
}

func parseArgs(ss []string) [][]int {
	args := make([][]int, 0, len(ss))
	for _, s := range ss {
		iss := readStrings(s, ",")
		arg := make([]int, 0, len(iss))
		for _, is := range iss {
			i, err := strconv.Atoi(is)
			if err != nil {
				panic(err)
			}
			arg = append(arg, i)
		}
		args = append(args, arg)
	}
	return args
}

func formatResults(results []string) string {
	return fmt.Sprintf("[%s]", strings.Join(results, ","))
}

type operation func(*LRUCache, []int) (*LRUCache, string)

var opMap = map[string]operation{
	`"LRUCache"`: func(_ *LRUCache, args []int) (*LRUCache, string) {
		obj := Constructor(args[0])
		return &obj, "null"
	},
	`"put"`: func(obj *LRUCache, args []int) (*LRUCache, string) {
		obj.Put(args[0], args[1])
		return obj, "null"
	},
	`"get"`: func(obj *LRUCache, args []int) (*LRUCache, string) {
		v := obj.Get(args[0])
		return obj, fmt.Sprintf("%d", v)
	},
}
