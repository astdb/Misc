/*
Given an arbitrary ransom note string and another string containing letters from all the magazines, write a function that will return true if the ransom note can be constructed from the magazines ; otherwise, it will return false.

Each letter in the magazine string can only be used once in your ransom note.



Example 1:

Input: ransomNote = "a", magazine = "b"
Output: false
Example 2:

Input: ransomNote = "aa", magazine = "ab"
Output: false
Example 3:

Input: ransomNote = "aa", magazine = "aab"
Output: true


Constraints:

You may assume that both strings contain only lowercase letters.
*/

package main

import (
	"log"
)

func main() {
	tests := [][]string{{"a", "b"}, {"aa", "ab"}, {"aa", "aab"}}

	for _, test := range tests {
		log.Printf("canConstruct(%v) == %v\n", test, canConstruct(test[0], test[1]))
	}
}

// check if each char of the note exists in the requisite number in the magazine string
func canConstruct(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}

	if magazine == ransomNote {
		return true
	}

	noteChars := map[rune]int{}
	for _, ch := range ransomNote {
		_, counted := noteChars[ch]

		if counted {
			noteChars[ch]++
		} else {
			noteChars[ch] = 1
		}
	}

	magChars := map[rune]int{}
	for _, ch := range magazine {
		_, counted := magChars[ch]

		if counted {
			magChars[ch]++
		} else {
			magChars[ch] = 1
		}
	}

	for noteCh, noteCount := range noteChars {
		magCount, exists := magChars[noteCh]

		if (!exists) || (magCount < noteCount) {
			return false
		}
	}

	return true
}
