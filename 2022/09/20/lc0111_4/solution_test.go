package main

import "testing"

func TestMinDepth(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "[3,9,20,null,null,15,7]",
			want:  2,
		},
		{
			input: "[2,null,3,null,4,null,5,null,6]",
			want:  5,
		},
		{
			input: "[]",
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			root := levelOrderRestore(parseStrings(tt.input))
			got := minDepth(root)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
