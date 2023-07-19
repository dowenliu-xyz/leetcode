package lc0049

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	tests := []struct {
		strs []string
		want [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{[]string{""}, [][]string{{""}}},
		{[]string{"a"}, [][]string{{"a"}}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := groupAnagrams(tt.strs)
			sortGroup(got)
			sortGroup(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func sortGroup(group [][]string) {
	for _, strs := range group {
		sort.Strings(strs)
	}
	sort.Slice(group, func(i, j int) bool {
		return group[i][0] < group[j][0]
	})
}
