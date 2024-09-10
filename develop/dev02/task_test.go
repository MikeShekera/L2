package main

import (
	"testing"
)

func TestUnpackStringBuilder(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{"Simple string", "a4", "aaaa", false},
		{"String with escape", `a4\5`, "aaaa5", false},
		{"String with double escape", `a4\\5`, `aaaa\\\\\`, false},
		{"String with escape and digit", `qwe\\5`, `qwe\\\\\`, false},
		{"Invalid string starting with digit", "4abc", "", true},
		{"String with escape at the end", `a\\`, "a\\", false},
		{"String with only escape", `\\`, `\`, false},
		{"Empty string", "", "", false},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := unpackStringBuilder(tt.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("unpackStringBuilder() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.expected {
					t.Errorf("unpackStringBuilder() = %v, expected %v", got, tt.expected)
				}
			},
		)
	}
}
