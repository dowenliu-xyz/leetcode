package lc0059

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_generateMatrix(t *testing.T) {
	tests := []struct {
		n    int
		want [][]int
	}{
		{1, [][]int{{1}}},
		{2, [][]int{{1, 2}, {4, 3}}},
		{3, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			got := generateMatrix(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
