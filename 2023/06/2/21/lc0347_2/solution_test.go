package lc0347

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
			want: []int{1, 2},
		},
		{
			nums: []int{1},
			k:    1,
			want: []int{1},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := topKFrequent(tt.nums, tt.k)
			sort.Ints(tt.want)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
