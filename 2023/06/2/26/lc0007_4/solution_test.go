package lc0007

import (
	"fmt"
	"testing"
)

func Test_reverse(t *testing.T) {
	tests := []struct {
		input  int
		expect int
	}{
		{input: 123, expect: 321},
		{input: -123, expect: -321},
		{input: 120, expect: 21},
		{input: 0, expect: 0},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if got := reverse(tt.input); got != tt.expect {
				t.Fatalf("expect %d, got %d", tt.expect, got)
			}
		})
	}
}
