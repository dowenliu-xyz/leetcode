package lc0889

import (
	"fmt"
	"strings"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/outputs"
)

func Test_constructFromPrePost(t *testing.T) {
	tests := []struct {
		preorder  []int
		postorder []int
		want      string
	}{
		{[]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1}, "[1,2,3,4,5,6,7]"},
		{[]int{1}, []int{1}, "[1]"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := constructFromPrePost(tt.preorder, tt.postorder)
			if !match(outputs.LevelOrderString(got), tt.want) {
				t.Errorf("constructFromPrePost(%v, %v) = %v, want %v", tt.preorder, tt.postorder, outputs.LevelOrderString(got), tt.want)
			}
		})
	}
}

func match(got, want string) bool {
	got = strings.ReplaceAll(got, " ", "")
	want = strings.ReplaceAll(want, " ", "")
	return got == want
}
