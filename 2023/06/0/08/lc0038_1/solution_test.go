package lc0038

import (
	"reflect"
	"testing"
)

func Test_permutation(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{
			input: `abc`,
			want:  []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
		{
			input: "aab",
			want:  []string{"aba", "aab", "baa"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := permutation(tt.input)
			if len(got) != len(tt.want) {
				t.Errorf("permutation(%s) = %s, want: %s", tt.input, got, tt.want)
			}
			want := stringSetOf(tt.want...)
			if got := stringSetOf(got...); !reflect.DeepEqual(got, want) {
				t.Errorf("permutation(%s) = %s, want: %s", tt.input, got, tt.want)
			}
		})
	}
}

func stringSetOf(strs ...string) map[string]struct{} {
	m := make(map[string]struct{})
	for i := range strs {
		m[strs[i]] = struct{}{}
	}
	return m
}
