package main

import (
	"fmt"
	"testing"
)

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		input1 string
		input2 string
		want   bool
	}{
		{
			input1: "[1,2,3]",
			input2: "[1,2,3]",
			want:   true,
		},
		{
			input1: "[1,2]",
			input2: "[1,null,2]",
			want:   false,
		},
		{
			input1: "[1,2,1]",
			input2: "[1,1,2]",
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s_%s", tt.input1, tt.input2), func(t *testing.T) {
			p := levelOrderRestore(parseStrings(tt.input1))
			q := levelOrderRestore(parseStrings(tt.input2))
			got := isSameTree(p, q)
			if tt.want != got {
				t.Errorf("want %t got %t", tt.want, got)
			}
		})
	}
}
