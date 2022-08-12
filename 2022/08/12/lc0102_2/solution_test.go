package lc0102_2

import "testing"

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "[3,9,20,null,null,15,7]",
			want:  "[[3],[9,20],[15,7]]",
		},
		{
			input: "[1]",
			want:  "[[1]]",
		},
		{
			input: "[]",
			want:  "[]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			root := levelOrderRestore(parseStrings(tt.input))
			result := levelOrder(root)
			got := formatResult(result)
			if tt.want != got {
				t.Errorf("want %s, got %s", tt.want, got)
			}
		})
	}
}
