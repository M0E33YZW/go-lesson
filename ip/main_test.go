package main

import (
	"strings"
	"testing"
)

func TestNewIpv4(t *testing.T) {
	tests := []struct {
		in      string
		want    Ipv4
		wantErr string
	}{
		{"172.16.254.1", Ipv4(0b10101100000100001111111000000001), ""},
		{"172.16.254.2", Ipv4(0b10101100000100001111111000000010), ""},
		{"-1.16.254.2", Ipv4(0), "0以上"},
		{"256.16.254.2", Ipv4(0), "255以下"},
	}

	for _, tt := range tests {
		have, err := NewIpv4(tt.in)
		if tt.wantErr != "" {
			if err == nil {
				t.Errorf("NewIpv4(%q): \n\t have: NoError \n\t want: Error", tt.in)
				continue
			}
			haveErr := err.Error()
			if !strings.Contains(haveErr, tt.wantErr) {
				t.Errorf("NewIpv4(%q): \n\t haveErr: %q \n\t wantErr: %q", tt.in, haveErr, tt.wantErr)
			}
			continue
		}
		if *have != tt.want {
			t.Errorf("NewIpv4(%q): \n\t have: %032b \n\t want: %032b", tt.in, *have, tt.want)
		}
	}
}

/*
Test実行
	go test .
Test実行・詳細表示
	go test -v .
*/
