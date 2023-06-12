package lc0020

import (
	"fmt"
	"testing"
)

func Test_isValid(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "()",
			want:  true,
		},
		{
			input: "()[]{}",
			want:  true,
		},
		{
			input: "(]",
			want:  false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if got := isValid(tt.input); got != tt.want {
				t.Errorf("expect %t, got %t", tt.want, got)
			}
		})
	}
}
