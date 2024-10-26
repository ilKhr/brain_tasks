package main

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Test",
			input:    "abba",
			expected: 2,
		},
		{
			name:     "Test",
			input:    "abcabcbb",
			expected: 3,
		},
		{
			name:     "Test",
			input:    "",
			expected: 0,
		},
		{
			name:     "Test",
			input:    "",
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := LengthOfLongestSubstring(test.input)

			if got != test.expected {
				t.Errorf("LengthOfLongestSubstring == %#v; want %#v", got, test.expected)
			}

		})
	}
}
