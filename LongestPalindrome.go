/*
Given a string which consists of lowercase or uppercase letters, find the length of the longest palindromes that can be built with those letters.

This is case sensitive, for example "Aa" is not considered a palindrome here.

Note:
Assume the length of given string will not exceed 1,010.

Example:

Input:
"abccccdd"

Output:
7

Explanation:
One longest palindrome that can be built is "dccaccd", whose length is 7.
*/

package main

import (
	"log"
)

func main() {
	tests := []string{"abccccdd"}

	for _, test := range tests {
		log.Printf("longestPalindrome(%s) == %d\n", test, longestPalindrome(test))
	}
}

// find how many characters occur twice or more in s
//
func longestPalindrome(s string) int {
	charCounts := map[rune]int{}
	ch2x := 0 // no of characters occuring twice or more in s
	for _, ch := range s {
		_, counted := charCounts[ch]
		if !counted {
			charCounts[ch] = 1
		} else {
			charCounts[ch]++
			ch2x++
		}
	}

	log.Println(charCounts)

	ch1x := false
	maxPalLen := 0
	for _, count := range charCounts {
		if count == 1 {
			// return ((ch2x * 2) + 1)
			ch1x = true
		} else if count/2 > 0 {
			maxPalLen += (count / 2) * 2

			if count%2 == 1 {
				ch1x = true
			}
		}
	}

	if ch1x {
		return maxPalLen + 1
	} else {
		return maxPalLen
	}
	// return (ch2x * 2)
}
