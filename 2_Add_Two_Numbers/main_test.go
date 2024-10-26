package main

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name     string
		inputl1  ListNode
		inputl2  ListNode
		expected *ListNode
	}{
		{
			name: "Test",
			inputl1: ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
			inputl2: ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			expected: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := AddTwoNumbers(&test.inputl1, &test.inputl2)

			isEqual := false

			expected := test.expected

			for expected != nil && got != nil {
				isEqual = expected.Val == got.Val

				got = got.Next
				expected = expected.Next
			}

			if !isEqual {
				t.Errorf("AddTwoNumbers == %#v; want %#v", got, test.expected)
			}

		})
	}
}
