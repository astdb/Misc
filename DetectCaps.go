/*
Given a word, you need to judge whether the usage of capitals in it is right or not.

We define the usage of capitals in a word to be right when one of the following cases holds:

All letters in this word are capitals, like "USA".
All letters in this word are not capitals, like "leetcode".
Only the first letter in this word is capital if it has more than one letter, like "Google".
Otherwise, we define that this word doesn't use capitals in a right way.
Example 1:
Input: "USA"
Output: True
Example 2:
Input: "FlaG"
Output: False
Note: The input will be a non-empty word consisting of uppercase and lowercase latin letters.
*/

package main

import (
	"fmt"
	"unicode"
)

func main() {
	tests := []string{"USA", "FlaG", "a", "A", ""}

	for _, test := range tests {
		fmt.Printf("detectCapitalUse(%s) == %v\n", test, detectCapitalUse(test))
	}

}

func detectCapitalUse(word string) bool {
	allUpper := true
	allLower := true

	for _, ch := range word {
		if unicode.IsLetter(ch) && unicode.IsLower(ch) {
			allUpper = false
		} else if unicode.IsLetter(ch) && !unicode.IsLower(ch) {
			allLower = false
		}
	}

	if allUpper || allLower {
		return true
	}

	wordRunes := []rune(word)
	if len(word) > 1 && unicode.IsLetter(wordRunes[0]) && !unicode.IsLower(wordRunes[0]) {
		allLower := true
		for j := 1; j < len(wordRunes); j++ {
			if !unicode.IsLower(wordRunes[j]) {
				return false
			}
		}

		return allLower
	}

	return false
}
