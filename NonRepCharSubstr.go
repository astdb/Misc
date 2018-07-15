/*
Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.


*/

package main

import (
	"fmt"
	"os"
)

func main() {
	// read input string from command line
	if len(os.Args) < 2 {
		fmt.Println("Please enter input string on command line.")
		return
	}

	inputStr := os.Args[1]
	fmt.Println(inputStr)
	fmt.Println(lengthOfLongestSubstring(inputStr))
}

func lengthOfLongestSubstring(s string) int {
	// transform string intu rune slice for easier manipulation/parsing
	inputSlice := []rune(s)

	// maximum length for the currently identified non-repeated character substring
	currentMaxLen := 0

	for i := 0; i < len(inputSlice); i++ {
		// start looking forward from i'th character forward for a non-repeated character substring
		thisSubStrLen := 0
		for j := i; j < len(inputSlice); j++ {
			// has this character been seen since s[i]?
			ch := inputSlice[j]
			charseen := false

			for k := i; k < j; k++ {
				if ch == inputSlice[k] {
					charseen = true;
				}
			}

			if charseen {
				break
			} else {
				thisSubStrLen++
			}
		}

		if thisSubStrLen > currentMaxLen {
			currentMaxLen = thisSubStrLen
		}
	}
	
	return currentMaxLen
}
