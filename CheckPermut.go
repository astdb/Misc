/*
Given two strings, check if one is a permutation of the other.

*/

package main

import (
	"log"
  "sort"
)

func main() {
	tests := [][]string{{"str1", "str2"}, {"str3", "str4"}, {"listen", "silent"}}

	for _, test := range tests {
		log.Printf("IsPerm(%s, %s) = %v\n", test[0], test[1], IsPerm(test[0], test[1]))
	}
}

func IsPerm(str1, str2 string) bool {
	if len(str1) != len(str2) {
		// permutations have to be of same length?
		return false
	}

	str1Runes := []rune(str1)
	str2Runes := []rune(str2)

	sort.Slice(str1Runes, func(i, j int) bool {
		return str1Runes[i] < str1Runes[j]
	})

	sort.Slice(str2Runes, func(i, j int) bool {
		return str2Runes[i] < str2Runes[j]
	})

	if string(str1Runes) == string(str2Runes) {
		return true
	}

	return false
}
