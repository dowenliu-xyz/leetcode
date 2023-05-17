package lc0015_1

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func Test_threeSum(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `[-1,0,1,2,-1,-4]`,
			want:  `[[-1,-1,2],[-1,0,1]]`,
		},
		{
			input: `[0,1,1]`,
			want:  `[]`,
		},
		{
			input: `[0,0,0]`,
			want:  `[[0,0,0]]`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			nums, err := inputs.ReadIntSlice(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			want, err := inputs.ReadIntSquareInLine(tt.want)
			if err != nil {
				t.Fatal(err)
			}
			got := threeSum(nums)
			matrixSort(got)
			matrixSort(want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("threeSum() = %v, want %v", got, want)
			}
		})
	}
}

func matrixSort(matrix [][]int) {
	for _, v := range matrix {
		sort.Ints(v)
	}
	sort.Slice(matrix, func(i, j int) bool {
		if matrix[i][0] != matrix[j][0] {
			return matrix[i][0] < matrix[j][0]
		}
		if matrix[i][1] != matrix[j][1] {
			return matrix[i][1] < matrix[j][1]
		}
		return matrix[i][2] < matrix[j][2]
	})
}
