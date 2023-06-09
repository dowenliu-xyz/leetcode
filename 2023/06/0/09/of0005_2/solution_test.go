package of0005

import (
	"fmt"
	"testing"
)

func Test_replaceSpace(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "We are happy.",
			want:  "We%20are%20happy.",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if got := replaceSpace(tt.input); got != tt.want {
				t.Errorf("replaceSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
