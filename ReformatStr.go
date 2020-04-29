/*
Given alphanumeric string s. (Alphanumeric string is a string consisting of lowercase English letters and digits).

You have to find a permutation of the string where no letter is followed by another letter and no digit is followed by another digit. That is, no two adjacent characters have the same type.

Return the reformatted string or return an empty string if it is impossible to reformat the string.



Example 1:

Input: s = "a0b1c2"
Output: "0a1b2c"
Explanation: No two adjacent characters have the same type in "0a1b2c". "a0b1c2", "0a1b2c", "0c2a1b" are also valid permutations.
Example 2:

Input: s = "leetcode"
Output: ""
Explanation: "leetcode" has only characters so we cannot separate them by digits.
Example 3:

Input: s = "1229857369"
Output: ""
Explanation: "1229857369" has only digits so we cannot separate them by characters.
Example 4:

Input: s = "covid2019"
Output: "c2o0v1i9d"
Example 5:

Input: s = "ab123"
Output: "1a2b3"
*/

package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	tests := []string{"", "a", "a1", "2", "a0b1c2", "leetcode", "1229857369", "covid2019", "ab123"}

	for testNo, test := range tests {
		log.Printf("#%d. reformat(%s) == %s\n", testNo, test, reformat(test))
	}
}

func reformat(s string) string {
	alphaCh := []rune{}
	numCh := []rune{}

	for _, ch := range s {
		if charNum(ch) {
			numCh = append(numCh, ch)
		} else {
			alphaCh = append(alphaCh, ch)
		}
	}

	if (len(alphaCh) == len(numCh)) || (abs(len(alphaCh)-len(numCh)) == 1) {
		var longer []rune
		var shorter []rune

		if len(alphaCh) >= len(numCh) {
			longer = alphaCh
			shorter = numCh
		} else {
			longer = numCh
			shorter = alphaCh
		}

		var strB strings.Builder
		appendLonger := true
		longIndex := 0
		shortIndex := 0
		for {
			if appendLonger {
				if longIndex < len(longer) {
					strB.WriteString(string(longer[longIndex]))
					longIndex++
					appendLonger = false
				}

			} else {
				if shortIndex < len(shorter) {
					strB.WriteString(string(shorter[shortIndex]))
					shortIndex++
					appendLonger = true
				}
			}

			if (longIndex >= len(longer)) && (shortIndex >= len(shorter)) {
				break
			}
		}

		return strB.String()

	}

	// cannot be reformatted
	return ""
}

func abs(x int) int {
	if x < 0 {
		return x * (-1)
	}

	return x
}

func charNum(ch rune) bool {
	_, err := strconv.Atoi(string(ch))
	if err != nil {
		return false
	}

	return true
}
