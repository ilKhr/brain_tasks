package main

import (
	"fmt"
)

type Index = int
type SliceElement = int

type DiffAIndex = int
type DifEl = int

func TwoSum(nums []int, target int) []int {
	cacheDif := make(map[DifEl]DiffAIndex)

	for index, num := range nums {
		dif := target - num

		if _, exists := cacheDif[dif]; exists {
			return []int{cacheDif[dif], index}
		} else {
			cacheDif[num] = index
		}
	}

	return []int{1, 1}
}

func main() {
	fmt.Println("Run: go test -v")
}
