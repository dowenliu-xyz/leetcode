package main

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{
			x:    123,
			want: 321,
		},
		{
			x:    -123,
			want: -321,
		},
		{
			x:    120,
			want: 21,
		},
		{
			x:    0,
			want: 0,
		},
		{
			x:    -2147483412,
			want: -2143847412,
		},
		{
			x:    1534236469,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.x), func(t *testing.T) {
			got := reverse(tt.x)
			if tt.want != got {
				t.Errorf("want %d, got %d", tt.want, got)
			}
		})
	}
}
