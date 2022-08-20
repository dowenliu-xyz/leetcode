package main

import "testing"

func TestInvertTree(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "[4,2,7,1,3,6,9]",
			want:  "[4,7,2,9,6,3,1]",
		},
		{
			input: "[2,1,3]",
			want:  "[2,3,1]",
		},
		{
			input: "[]",
			want:  "[]",
		},
		{
			input: "[4,2,7,null,3,6,9]",
			want:  "[4,7,2,9,6,3]",
		},
		{
			input: "[4,2,7,null,3,6]",
			want:  "[4,7,2,null,6,3]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			root := levelOrderRestore(parseStrings(tt.input))
			got := levelOrderFormat(invertTree(root))
			if tt.want != got {
				t.Errorf("want %s, got %s", tt.want, got)
			}
		})
	}
}
