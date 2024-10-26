package main

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			arr    []int
			target int
		}
		expected []int
	}{
		{
			name: "Test",
			input: struct {
				arr    []int
				target int
			}{
				arr:    []int{2, 7, 11, 15},
				target: 9,
			},
			expected: []int{0, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := TwoSum(test.input.arr, test.input.target)

			if !slices.Equal(got, test.expected) {
				t.Errorf("TwoSum = %#v; want %#v", got, test.expected)
			}

		})
	}
}
