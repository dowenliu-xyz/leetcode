package main

import "testing"

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "[3,9,20,null,null,15,7]",
			want:  3,
		},
		{
			input: "[]",
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			root := levelOrderRestore(parseStrings(tt.input))
			got := maxDepth(root)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
