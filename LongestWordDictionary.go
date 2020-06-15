/*
Given a list of strings words representing an English Dictionary, find the longest word in words that can be built one character at a time by other words in words. If there is more than one possible answer, return the longest word with the smallest lexicographical order.

If there is no answer, return the empty string.
Example 1:
Input:
words = ["w","wo","wor","worl", "world"]
Output: "world"
Explanation:
The word "world" can be built one character at a time by "w", "wo", "wor", and "worl".
Example 2:
Input:
words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
Output: "apple"
Explanation:
Both "apply" and "apple" can be built from other words in the dictionary. However, "apple" is lexicographically smaller than "apply".
Note:

All the strings in the input will only contain lowercase letters.
The length of words will be in the range [1, 1000].
The length of words[i] will be in the range [1, 30].
*/

package main

import (
	"log"
	"sort"
)

func main() {
	tests := [][]string{{"w", "wo", "wor", "worl", "world"}, {"a", "banana", "app", "appl", "ap", "apply", "apple"}}

	for _, test := range tests {
		log.Printf("longestWords(%v) == %s\n", test, longestWord(test))
	}
}

func longestWord(words []string) string {
	// build word index
	wordMap := map[string]int{}

	for _, word := range words {
		wordMap[word] = 1
	}

	// log.Println(wordMap)

	longest := 0
	res := []string{}
	for _, word := range words {
		// if eligible(word, wordMap) && len(word) >= longest {
		if eligible(word, wordMap) {
			// res = word
			if len(word) > longest {
				res = []string{}
				res = append(res, word)
				longest = len(word)
			} else if len(word) == longest {
				res = append(res, word)
			}
		}
	}

	if len(res) == 0 {
		return ""
	}

	if len(res) == 1 {
		return res[0]
	}

	sort.Strings(res)
	return res[0]
}

func eligible(word string, wordMap map[string]int) bool {
	for i := 1; i <= len(word); i++ {
		thisSubStr := word[:i]

		_, wordFound := wordMap[thisSubStr]
		if !wordFound {
			return false
		}
	}

	return true
}
