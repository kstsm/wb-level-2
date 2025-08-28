package main

import "testing"

func TestUnpackString(t *testing.T) {
	tests := []struct {
		str      string
		expected string
		wantErr  bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"", "", false},
		{"qwe\\4\\5", "qwe45", false},
		{"qwe\\45", "qwe44444", false},
	}

	for _, tt := range tests {
		got, err := unpackString(tt.str)
		if (err != nil) != tt.wantErr {
			t.Errorf("unpackString(%q) error = %v, wantErr %v", tt.str, err, tt.wantErr)
			continue
		}
		if got != tt.expected {
			t.Errorf("unpackString(%q) = %q, want %q", tt.str, got, tt.expected)
		}
	}
}
