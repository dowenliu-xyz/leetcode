package lc0072

import (
	"fmt"
	"testing"
)

func Test_minDistance(t *testing.T) {
	tests := []struct {
		word1 string
		word2 string
		want  int
	}{
		{"horse", "ros", 3},
		{"intention", "execution", 5},
		{"", "", 0},
		{"distance", "springbok", 9},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if got := minDistance(tt.word1, tt.word2); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
