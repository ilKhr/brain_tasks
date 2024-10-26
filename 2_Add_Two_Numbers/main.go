package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeNode(val int, prevListNode *ListNode) *ListNode {
	newNode := &ListNode{
		Val:  val,
		Next: nil,
	}

	if prevListNode != nil {
		prevListNode.Next = newNode
	}

	return newNode
}

func sumTwoNumber(a int, b int, c int) (main int, additional int) {
	if a > 9 || b > 9 || c > 9 {
		panic("Number can't be more than 9")
	}

	sum := a + b + c

	main = sum % 10

	sum = sum / 10

	additional = sum % 10

	return main, additional
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var header *ListNode
	var current *ListNode
	additionalOne := 0
	main := 0
	additional := 0

	for l1 != nil || l2 != nil {
		fmt.Println("additionalOne", additionalOne)
		if l1 != nil && l2 != nil {
			main, additional = sumTwoNumber(l1.Val, l2.Val, additionalOne)
		}

		if l1 == nil && l2 != nil {
			main, additional = sumTwoNumber(l2.Val, 0, additionalOne)
		}

		if l1 != nil && l2 == nil {
			main, additional = sumTwoNumber(l1.Val, 0, additionalOne)
		}

		if l1 == nil && l2 == nil && additionalOne != 0 {
			main, additional = sumTwoNumber(0, 0, additionalOne)
		}

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

		current = makeNode(main, current)

		additionalOne = additional

		if header == nil {
			header = current
		}
	}

	if additional != 0 {
		makeNode(additional, current)
	}

	return header
}

func main() {
	fmt.Println("Run: go test -v")
}
