package main

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1   string
		l2   string
		want string
	}{
		{
			l1:   "[2,4,3]",
			l2:   "[5,6,4]",
			want: "[7,0,8]",
		},
		{
			l1:   "[0]",
			l2:   "[0]",
			want: "[0]",
		},
		{
			l1:   "[9,9,9,9,9,9,9]",
			l2:   "[9,9,9,9]",
			want: "[8,9,9,9,0,0,0,1]",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s + %s", tt.l1, tt.l2), func(t *testing.T) {
			l1, l2 := readInput(tt.l1), readInput(tt.l2)
			add := addTwoNumbers(l1, l2)
			got := formatList(add)
			if tt.want != got {
				t.Errorf("want %s, got %s", tt.want, got)
			}
		})
	}
}
