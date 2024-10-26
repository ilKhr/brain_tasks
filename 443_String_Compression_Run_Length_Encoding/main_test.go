package main

import (
	"bytes"
	"testing"
)

func TestCompress(t *testing.T) {
	tests := []struct {
		name           string
		input          []byte
		expectedLength int
		expectedInput  []byte
	}{
		{
			name:           "Test",
			input:          []byte{'a', 'b', 'c', 'c', 'c', 'c', 'c', 'c'},
			expectedLength: 4,
			expectedInput:  []byte{'a', 'b', 'c', '6', 'c', 'c', 'c', 'c'},
		},
		{
			name:           "Test",
			input:          []byte{'a'},
			expectedLength: 1,
			expectedInput:  []byte{'a'},
		},
		{
			name:           "Test",
			input:          []byte{'a', 'a', 'a'},
			expectedLength: 2,
			expectedInput:  []byte{'a', '3', 'a'},
		},
		{
			name:           "Test",
			input:          []byte{},
			expectedLength: 0,
			expectedInput:  []byte{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			input := make([]byte, len(test.input))
			copy(input, test.input)

			got := CompressWithPointers(input)

			if got != test.expectedLength || !(bytes.Equal(input, test.expectedInput)) {
				t.Errorf("CompressWithPointers(input) = %d; want %d", got, test.expectedLength)
				t.Errorf("CompressWithPointers: input; want %#v, got %#v", test.expectedInput, input)
			}

		})
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			input := make([]byte, len(test.input))
			copy(input, test.input)

			got := CompressWithEndSymbol(input)

			if got != test.expectedLength || !(bytes.Equal(input, test.expectedInput)) {
				t.Errorf("CompressWithEndSymbol(input) = %d; want %d", got, test.expectedLength)
				t.Errorf("CompressWithEndSymbol: input; want %#v, got %#v", test.expectedInput, input)
			}

		})
	}

}
