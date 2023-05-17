package lc0013_1

import (
	"fmt"
	"testing"
)

func Test_romanToInt(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if got := romanToInt(tt.input); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
