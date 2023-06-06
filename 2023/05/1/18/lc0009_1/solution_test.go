package lc0009_1

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{input: 121, want: true},
		{input: -121, want: false},
		{input: 10, want: false},
		{input: -101, want: false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if got := isPalindrome(tt.input); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
