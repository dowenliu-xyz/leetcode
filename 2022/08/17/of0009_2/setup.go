package main

import (
	"fmt"
	"strconv"
	"strings"
)

func readStrings(input string) []string {
	input = strings.TrimSpace(input)
	if strings.HasPrefix(input, "[") {
		input = input[1:]
	}
	if strings.HasSuffix(input, "]") {
		input = input[:len(input)-1]
	}
	parts := strings.Split(input, ",")
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
		iss := readStrings(s)
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

type operation func(*CQueue, []int) (*CQueue, string)

var opMap = map[string]operation{
	`"CQueue"`: func(_ *CQueue, _ []int) (*CQueue, string) {
		obj := Constructor()
		return &obj, "null"
	},
	`"appendTail"`: func(obj *CQueue, args []int) (*CQueue, string) {
		obj.AppendTail(args[0])
		return obj, "null"
	},
	`"deleteHead"`: func(obj *CQueue, _ []int) (*CQueue, string) {
		v := obj.DeleteHead()
		return obj, fmt.Sprintf("%d", v)
	},
}
