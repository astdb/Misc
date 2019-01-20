/*
We are given two sentences A and B.  (A sentence is a string of space separated words.  Each word consists only of lowercase letters.)

A word is uncommon if it appears exactly once in one of the sentences, and does not appear in the other sentence.

Return a list of all uncommon words. 

You may return the list in any order. 

Example 1:

Input: A = "this apple is sweet", B = "this apple is sour"
Output: ["sweet","sour"]
Example 2:

Input: A = "apple apple", B = "banana"
Output: ["banana"]
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	tests := [][]string{{"this apple is sweet", "this apple is sour"}, {"apple apple", "banana"}}

	for _, test := range tests {
		fmt.Println(uncommonFromSentences(test[0], test[1]))
	}
}

func uncommonFromSentences(strA string, strB string) []string {
	strAArr := strings.Split(strA, " ")
	strBArr := strings.Split(strB, " ")

	strAWordMap := map[string]int{}
	strBWordMap := map[string]int{}

	for _, word := range strAArr {
		_, exists := strAWordMap[word]

		if exists {
			strAWordMap[word]++
		} else {
			strAWordMap[word] = 1
		}
	}

	for _, word := range strBArr {
		_, exists := strBWordMap[word]

		if exists {
			strBWordMap[word]++
		} else {
			strBWordMap[word] = 1
		}
	}

	uncommonWords := []string{}
	for word, count := range strAWordMap {
		if count == 1 {
			_, exists := strBWordMap[word]
			if !exists {
				uncommonWords = append(uncommonWords, word)
			}
		}
	}

	for word, count := range strBWordMap {
		if count == 1 {
			_, exists := strAWordMap[word]
			if !exists {
				uncommonWords = append(uncommonWords, word)
			}
		}
	}

	return uncommonWords
}
