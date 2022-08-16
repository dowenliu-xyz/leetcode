package main

import (
	"fmt"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		input      string
		p, q, want int
	}{
		{
			input: "[3,5,1,6,2,0,8,null,null,7,4]",
			p:     5,
			q:     1,
			want:  3,
		},
		{
			input: "[3,5,1,6,2,0,8,null,null,7,4]",
			p:     5,
			q:     4,
			want:  5,
		},
		{
			input: "[1,2]",
			p:     1,
			q:     2,
			want:  1,
		},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%s_%d_%d", tt.input, tt.p, tt.q)
		t.Run(name, func(t *testing.T) {
			root := levelOrderRestore(parseStrings(tt.input))
			p := nodeOfVal(root, tt.p)
			q := nodeOfVal(root, tt.q)
			got := lowestCommonAncestor(root, p, q).Val
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
