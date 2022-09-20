package main

import "testing"

func TestIsCompleteTree(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "[1,2,3,4,5,6]",
			want:  true,
		},
		{
			input: "[1,2,3,4,5,null,7]",
			want:  false,
		},
		{
			input: "[]",
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			root := levelOrderRestore(parseStrings(tt.input))
			got := isCompleteTree(root)
			if tt.want != got {
				t.Errorf("want %t, got %t", tt.want, got)
			}
		})
	}
}
