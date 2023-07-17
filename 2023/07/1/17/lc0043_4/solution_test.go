package lc0043

import "testing"

func Test_multiply(t *testing.T) {
	tests := []struct {
		num1 string
		num2 string
		want string
	}{
		{"2", "3", "6"},
		{"123", "456", "56088"},
	}
	for _, tt := range tests {
		t.Run(tt.num1+"*"+tt.num2, func(t *testing.T) {
			if got := multiply(tt.num1, tt.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
