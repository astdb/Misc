/*
Check if a given string consists of all unique characters.

What if external datastructures cannot be used?


NOTES:
01. Assume string consists of Unicode chars. Use a rune->bool map to keep track of characters observed in string.

*/

package main

import (
	"log"
)

func main() {
	tests := []string{"A", "AB", "ABC", "ABBC"}

	for _, test := range tests {
		log.Printf("strUniqueCh(%s) = %v\n", test, strUniqueCh(test))
	}
}

// return true if input string consists of unique characters: else false.
func strUniqueCh(str string) bool {
	// map with string's chars as keys
	charMap := map[rune]bool{}

	// transform string to Unicode character array
	strChArr := []rune(str)

	for _, ch := range strChArr {
		_, charSeen := charMap[ch]

		if charSeen {
			// duplicate occurence
			return false

		} else {
			charMap[ch] = true

		}
	}

	return true
}
