/*
Given a string, you need to reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

Example 1:
Input: "Let's take LeetCode contest"
Output: "s'teL ekat edoCteeL tsetnoc"
Note: In the string, each word is separated by single space and there will not be any extra space in the string.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}

func reverseWords(s string) string {
	// transform input to rune slice (for character manipulation)
	s_runes := []rune(strings.TrimSpace(s))

	// starting index of the last unreversed word
	start := 0

	// iterate through input
	for i := 0; i < len(s_runes); i++ {

		// if the current character is a whitespace or end of string reached, reverse the current word
		if s_runes[i] == ' ' || i == len(s_runes)-1 {

			// if the end of string is reached, increment i by one to ensure the current word is fully reversed including the last character
			if i == len(s_runes)-1 {
				i++
			}

			wordLen := i - start // length of current word
			var k int            // variable to hold a decreasing pointer from the end index of the current word

			// reverse chars
			for j := start; j < (start + (wordLen / 2)); j++ {
				if j == start {
					k = j + wordLen - 1
				} else {
					k--
				}

				s_runes[j], s_runes[k] = s_runes[k], s_runes[j]
			}

			// reset start to the beginning of the next word
			start = i + 1
		}
	}

	return string(s_runes)
}
