package main

import (
	"fmt"
	"strconv"
)

func CompressWithPointers(chars []byte) int {
	if len(chars) <= 1 {
		return len(chars)
	}

	writePointer := 0
	startPointer := 0
	finishPointer := 0

	for finishPointer <= len(chars) {
		if finishPointer < len(chars) && chars[startPointer] == chars[finishPointer] {
			finishPointer++
		} else {
			chars[writePointer] = chars[startPointer]
			writePointer++

			if finishPointer-startPointer > 1 {

				for index := 0; index < len([]byte(strconv.Itoa(finishPointer-startPointer))); index++ {
					chars[writePointer] = []byte(strconv.Itoa(finishPointer - startPointer))[index]
					writePointer++
				}
			}

			if finishPointer < len(chars) {
				startPointer = finishPointer
			} else {
				finishPointer++
			}
		}
	}

	return writePointer
}

func CompressWithEndSymbol(chars []byte) int {
	if len(chars) == 0 {
		return len(chars)
	}

	s := chars[:0]
	var count int

	tmpChars := append(chars, '\n')

	for i := 0; tmpChars[i] != '\n'; i++ {
		isSequenceEnd := tmpChars[i] != tmpChars[i+1]

		count++

		if isSequenceEnd {
			s = append(s, tmpChars[i])

			if count > 1 {
				s = append(s, []byte(strconv.Itoa(count))...)
			}

			count = 0
		}

	}

	fmt.Printf("S: %#v \n", s)

	return len(s)
}

func main() {
	fmt.Println("Run: go test -v")
}
