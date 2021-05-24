/*
Given two strings a and b, return the length of the longest uncommon subsequence between a and b. If the longest uncommon subsequence does not exist, return -1.

An uncommon subsequence between two strings is a string that is a subsequence of one but not the other.

A subsequence of a string s is a string that can be obtained after deleting any number of characters from s.

For example, "abc" is a subsequence of "aebdc" because you can delete the underlined characters in "aebdc" to get "abc". Other subsequences of "aebdc" include "aebdc", "aeb", and "" (empty string).


Example 1:

Input: a = "aba", b = "cdc"
Output: 3
Explanation: One longest uncommon subsequence is "aba" because "aba" is a subsequence of "aba" but not "cdc".
Note that "cdc" is also a longest uncommon subsequence.
Example 2:

Input: a = "aaa", b = "bbb"
Output: 3
Explanation: The longest uncommon subsequences are "aaa" and "bbb".
Example 3:

Input: a = "aaa", b = "aaa"
Output: -1
Explanation: Every subsequence of string a is also a subsequence of string b. Similarly, every subsequence of string b is also a subsequence of string a.


Constraints:

1 <= a.length, b.length <= 100
a and b consist of lower-case English letters.

*/

package main

import (
	"log"
	"strings"
)

func main() {
	tests := [][]string{{"aba", "cdc"}, {"aaa", "bbb"}, {"aaa", "aaa"}}
	for _, test := range tests {
		log.Printf("findLUSlength(%s, %s) == %d\n", test[0], test[1], findLUSlength(test[0], test[1]))
	}
}

// per substring of a, if its a substring of b
// if yes, ignore. if not, check if length > current longest uncommon substr - update if longer
// return longest uncommon subsr length
// optimization: start searching from longest possible substrings (e.g. len(b) if len(b) < len(a), else len(a))
func findLUSlength(a string, b string) int {
	var subStrLen int
	if len(b) < len(a) {
		subStrLen = len(b)
	} else {
		subStrLen = len(a)
	}

	var longestSubstrLen int = -1

	for subStrLen > 0 {
		// get all substrings of a of subStrlen
		subStrs := getAllSubStrs(a, subStrLen)

		for _, subStr := range subStrs {
			if !strings.Contains(b, subStr) {
				if len(subStr) > longestSubstrLen {
					longestSubstrLen = len(subStr)
				}
			}
		}

		subStrLen--
	}

	return longestSubstrLen
}

func getAllSubStrs(str string, strLen int) []string {
	subStrs := []string{}

	for start := 0; (start + strLen) <= len(str); start++ {
		end := start + strLen
		subStrs = append(subStrs, str[start:end])
	}

	return subStrs
}
