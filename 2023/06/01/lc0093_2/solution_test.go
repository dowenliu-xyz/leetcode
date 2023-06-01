package lc0093_2

import (
	"fmt"
	"testing"
)

func Test_restoreIpAddresses(t *testing.T) {
	tests := []struct {
		s    string
		want []string
	}{
		{
			s:    "25525511135",
			want: []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			s:    "0000",
			want: []string{"0.0.0.0"},
		},
		{
			s:    "1",
			want: []string{},
		},
		{
			s:    "101023",
			want: []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"},
		},
		{
			s:    "19216811",
			want: []string{"1.92.168.11", "19.2.168.11", "19.21.68.11", "19.216.8.11", "19.216.81.1", "192.1.68.11", "192.16.8.11", "192.16.81.1", "192.168.1.1"},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := restoreIpAddresses(tt.s)
			if !stringSetEqual(got, tt.want) {
				t.Errorf("restoreIpAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func stringSetEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[string]struct{}, len(a))
	for _, v := range a {
		m[v] = struct{}{}
	}
	for _, v := range b {
		if _, ok := m[v]; !ok {
			return false
		}
	}
	return true
}
