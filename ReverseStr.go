/*
Write a function that takes a string as input and returns the string reversed.

Example:
Given s = "hello", return "olleh".
*/

package main

import (
	"fmt"
)

func main() {
	tests := [][]string{{"", ""}, {" ", " "}, {"   ", "   "}, {"a", "a"}, {"ab", "ba"}, {"hello", "olleh"}, {"Hello, 世界!", "!界世 ,olleH"}}

	for _, test := range tests {
		if reverseString(test[0]) == test[1] {
			fmt.Println("PASS")
		} else {
			fmt.Println("FAIL")
			fmt.Println("\t", test)
		}
	}
}

func reverseString(s string) string {
	// transform input string into rune (unicode char) slice for convenient manipulation
	s_runes := []rune(s)

	// reverse characters
	for i := 0; i < len(s_runes)/2; i++ {
		temp := s_runes[i]
		s_runes[i] = s_runes[len(s_runes)-i-1]
		s_runes[len(s_runes)-i-1] = temp
	}

	return string(s_runes)
}
