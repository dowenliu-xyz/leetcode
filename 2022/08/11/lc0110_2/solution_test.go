package lc0110_2

import "testing"

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "[null]",
			want:  true,
		},
		{
			input: "[3,9,20,null,null,15,7]",
			want:  true,
		},
		{
			input: "[1,2,2,3,3,null,null,4,4]",
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
			got := isBalanced(root)
			if tt.want != got {
				t.Errorf("want %t, got %t", tt.want, got)
			}
		})
	}
}
