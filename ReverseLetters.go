/*
Given a string S, return the "reversed" string where all characters that are not a letter stay in the same place, and all letters reverse their positions.

Example 1:
Input: "ab-cd"
Output: "dc-ba"


Example 2:
Input: "a-bC-dEf-ghIj"
Output: "j-Ih-gfE-dCba"


Example 3:
Input: "Test1ng-Leet=code-Q!"
Output: "Qedo1ct-eeLg=ntse-T!"

*/

package main

import (
	"fmt"
	"unicode"
)

func main() {
	tests := []string{"", "a", "abc", "a-bcd", "ab-cd"}

	for _, test := range tests {
		fmt.Println(test, "=>", reverseOnlyLetters2(test))
	}
}

// O(1) space
func reverseOnlyLetters2(S string) string {
	// convert string to slice of runes
	s_runes := []rune(S)

	// swap letters from start and end
	i := 0                // start index
	j := len(s_runes) - 1 // end index

	for i < j {
		// if both are letters, swap
		if unicode.IsLetter(s_runes[i]) && unicode.IsLetter(s_runes[j]) {
			tmp := s_runes[i]
			s_runes[i] = s_runes[j]
			s_runes[j] = tmp

			i++
			j--
		} else {
			// one or both of the letters not letters
			if !unicode.IsLetter(s_runes[i]) {
				i++
			}

			if !unicode.IsLetter(s_runes[j]) {
				j--
			}
		}
	}

	return string(s_runes)
}

// O(n) space
func reverseOnlyLetters(S string) string {
	s_runes := []rune(S)

	// compile list of letters in S in reverse order
	sLettersReverse := []rune{}
	for i := len(s_runes) - 1; i >= 0; i-- {
		if unicode.IsLetter(s_runes[i]) {
			sLettersReverse = append(sLettersReverse, s_runes[i])
		}
	}

	// traverse s from start, replacing letters with elements on sLettersReverse from start
	j := 0 // index for sLettersReverse
	for i := 0; i < len(s_runes); i++ {
		if unicode.IsLetter(s_runes[i]) {
			s_runes[i] = sLettersReverse[j]
			j++
		}
	}

	return string(s_runes)
}
