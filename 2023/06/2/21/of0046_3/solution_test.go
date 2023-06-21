package of0046

import (
	"fmt"
	"testing"
)

func Test_translateNum(t *testing.T) {
	tests := []struct {
		num  int
		want int
	}{
		{12258, 5},
		{25, 2},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if got := translateNum(tt.num); got != tt.want {
				t.Errorf("translateNum(%d) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}
