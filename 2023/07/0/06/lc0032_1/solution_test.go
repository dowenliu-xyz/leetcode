package lc0032

import (
	"fmt"
	"testing"
)

func Test_longestValidParentheses(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"(()", 2},
		{")()())", 4},
		{"", 0},
		{"(", 0},
		{")", 0},
		{"()(()", 2},
		{")()())()())", 4},
		{"((()()))", 8},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if got := longestValidParentheses(tt.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
