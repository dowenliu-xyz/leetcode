package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		input int
		want  []string
	}{
		{
			input: 3,
			want:  []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			input: 1,
			want:  []string{"()"},
		},
		{
			input: 0,
			want:  []string{},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.input), func(t *testing.T) {
			got := generateParenthesis(tt.input)
			sort.Strings(tt.want)
			sort.Strings(got)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("want %q, got %q", tt.want, got)
			}
		})
	}
}
