package lc0017

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{
			input: "23",
			want:  []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			input: "",
			want:  []string{},
		},
		{
			input: "2",
			want:  []string{"a", "b", "c"},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := letterCombinations(tt.input)
			sort.Strings(got)
			sort.Strings(tt.want)
			if len(tt.want) == 0 && len(got) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
