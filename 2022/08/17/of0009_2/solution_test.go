package main

import (
	"fmt"
	"testing"
)

func TestCQueue(t *testing.T) {
	tests := []struct {
		input1 string
		input2 string
		want   string
	}{
		{
			input1: `["CQueue","appendTail","deleteHead","deleteHead"]`,
			input2: "[[],[3],[],[]]",
			want:   "[null,null,3,-1]",
		},
		{
			input1: `["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]`,
			input2: "[[],[],[5],[2],[],[]]",
			want:   "[null,-1,null,null,5,2]",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ops := readStrings(tt.input1)
			argsList := parseArgs(readStrings(tt.input2))
			var results []string
			var obj *CQueue
			for j, op := range ops {
				var result string
				obj, result = opMap[op](obj, argsList[j])
				results = append(results, result)
			}
			got := formatResults(results)
			if tt.want != got {
				t.Errorf("want %s, got %s", tt.want, got)
			}
		})
	}
}
