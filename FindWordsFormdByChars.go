/*
You are given an array of strings words and a string chars.

A string is good if it can be formed by characters from chars (each character can only be used once).

Return the sum of lengths of all good strings in words.



Example 1:

Input: words = ["cat","bt","hat","tree"], chars = "atach"
Output: 6
Explanation:
The strings that can be formed are "cat" and "hat" so the answer is 3 + 3 = 6.

Example 2:

Input: words = ["hello","world","leetcode"], chars = "welldonehoneyr"
Output: 10
Explanation:
The strings that can be formed are "hello" and "world" so the answer is 5 + 5 = 10.


Note:

1 <= words.length <= 1000
1 <= words[i].length, chars.length <= 100
All strings contain lowercase English letters only.

*/

package main

import (
	"log"
)

func main() {
	tests := [][][]string{{{"cat", "bt", "hat", "tree"}, {"atach"}}, {{"hello", "world", "leetcode"}, {"welldonehoneyr"}}}

	for _, test := range tests {
		log.Printf("countCharacters(%v, %s) == %d\n", test[0], test[1][0], countCharacters(test[0], test[1][0]))
	}
}

func countCharacters(words []string, chars string) int {
	// create map of chars (with chars as keys) of the chars string
	charMap := map[rune]int{}

	for _, ch := range chars {
		_, inMap := charMap[ch]
		if inMap {
			charMap[ch]++
		} else {
			charMap[ch] = 1
		}
	}

	// check for each word given, if it can be constructed using given char set
	goodCount := 0

	for _, word := range words {
		mapCopy := map[rune]int{}
		for k, v := range charMap {
			mapCopy[k] = v
		}

		if isGood(word, mapCopy) {
			goodCount += len(word)
		}
	}

	return goodCount
}

// check if every char in a given string exist in a given char map
func isGood(str string, charsMap map[rune]int) bool {
	for _, ch := range str {
		_, chExists := charsMap[ch]
		if chExists {
			if charsMap[ch] > 0 {
				charsMap[ch]--
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
