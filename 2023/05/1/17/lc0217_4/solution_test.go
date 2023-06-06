package lc0217_4

import (
	"fmt"
	"testing"

	"github.com/dowenliu-xyz/leetcode/internal/inputs"
)

func Test_containsDuplicate(t *testing.T) {
	tests := []struct {
		input  string
		expect bool
	}{
		{"[]", false},
		{"[1]", false},
		{"[1,2]", false},
		{"[1,1]", true},
		{"[1,2,3,1]", true},
		{"[1,2,3,4]", false},
		{"[1,1,1,3,3,4,3,2,4,2]", true},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			nums, err := inputs.ReadIntSlice(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			if got := containsDuplicate(nums); got != tt.expect {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.expect)
			}
		})
	}
}
