/*
In an alien language, surprisingly they also use english lowercase letters, but possibly in a different order. The order of the alphabet is some permutation of lowercase letters.

Given a sequence of words written in the alien language, and the order of the alphabet, return true if and only if the given words are sorted lexicographicaly in this alien language.

 

Example 1:

Input: words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
Output: true
Explanation: As 'h' comes before 'l' in this language, then the sequence is sorted.
Example 2:

Input: words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz"
Output: false
Explanation: As 'd' comes after 'l' in this language, then words[0] > words[1], hence the sequence is unsorted.
Example 3:

Input: words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz"
Output: false
Explanation: The first three characters "app" match, and the second string is shorter (in size.) According to lexicographical rules "apple" > "app", because 'l' > '∅', where '∅' is defined as the blank character which is less than any other character (More info).
 

Note:

1 <= words.length <= 100
1 <= words[i].length <= 20
order.length == 26
All characters in words[i] and order are english lowercase letters.

*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(isAlienSorted([]string{"hello","leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"))
	fmt.Println(isAlienSorted([]string{"word","world","row"}, "worldabcefghijkmnpqstuvxyz"))
	fmt.Println(isAlienSorted([]string{"apple","app"}, "abcdefghijklmnopqrstuvwxyz"))
}

func isAlienSorted(words []string, order string) bool {
	letterIndexInvalid := 0
	letterIndex := 0
	inOrder := false

	for letterIndexInvalid < len(words) && !inOrder {
		letterIndexInvalid = 0
		inOrder = true

		// for each word, check letterIndex/letterIndex+1 chars are in order
		for wordIndex, word := range words {
			wordRunes := []rune(word)

			if letterIndex < len(wordRunes) {
				// in == true
				// letterIndexInvalid = false
				if wordIndex+1 < len(words) {

					if letterIndex < len([]rune(words[wordIndex+1])) {
						ord := alphabetOrder([]rune(words[wordIndex])[letterIndex], []rune(words[wordIndex+1])[letterIndex], []rune(order))

						if ord < 0 {
							// out of order
							return false
						}
	
						if ord == 0 {
							// same order, need to check next index
							inOrder = false
						}
					}
					
				}

			} else {
				letterIndexInvalid++
			}	
	
		}

		letterIndex++
	}

	return true
}

// given two letters return:
// 	- 0 if they're in the same alphabetical order in the given alphabet
//	- 1 if the first letter is before of the second in the given alphabet
//	- -1 if the first letter is after the second in the given alphabet
func alphabetOrder(ch1, ch2 rune, alphabet []rune) int {
	for _, thisCh := range alphabet {
		if ch1 == thisCh && ch2 == thisCh {
			return 0
		}

		if ch1 == thisCh {
			return 1
		}

		if ch2 == thisCh {
			return -1
		}
	}

	return 3
}
