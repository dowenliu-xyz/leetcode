package lc0101_1

import "testing"

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "[1,2,2,3,4,4,3]",
			want:  true,
		},
		{
			input: "[]",
			want:  true,
		},
		{
			input: "[1,2,2,null,3,null,3]",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			root := levelOrderRestore(parseStrings(tt.input))
			got := isSymmetric(root)
			if tt.want != got {
				t.Errorf("want %t, got %t", tt.want, got)
			}
		})
	}
}
