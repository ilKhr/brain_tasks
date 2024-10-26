package main

import (
	"fmt"
)

type Symbol = rune
type SymbolIndex = int

type ProcessMap = map[Symbol]SymbolIndex

func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	processMap := make(ProcessMap)

	maxLength := 0
	startPointer := 0

	for index, symbol := range s {

		if item, exists := processMap[symbol]; exists {

			if item >= startPointer {
				startPointer = item + 1
			}

			delete(processMap, symbol)
		}

		processMap[symbol] = index

		currentLength := index - startPointer

		if currentLength > maxLength {

			maxLength = currentLength

		}

	}

	return maxLength + 1
}

func main() {
	fmt.Println("Run: go test -v")
}
