package lc0022

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{input: 1, want: `["()"]`},
		{input: 2, want: `["(())","()()"]`},
		{input: 3, want: `["((()))","(()())","(())()","()(())","()()()"]`},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := generateParenthesis(tt.input)
			if got := sprintStrings(got); got != tt.want {
				t.Errorf("generateParenthesis(%d) = %s, want %s", tt.input, got, tt.want)
			}
		})
	}
}

func sprintStrings(ss []string) string {
	buf := bytes.Buffer{}
	buf.WriteByte('[')
	for i, s := range ss {
		if i > 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(strconv.Quote(s))
	}
	buf.WriteByte(']')
	return buf.String()
}
