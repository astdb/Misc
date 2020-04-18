/*
Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.

Examples:

s = "leetcode"
return 0.

s = "loveleetcode",
return 2.
Note: You may assume the string contain only lowercase letters.
*/

package main

import (
	"log"
)

func main() {
	tests := []string{"", "123", "a", "aa", "leetcode", "loveleetcode"}

	for _, test := range tests {
		log.Printf("firstUniqChar(%q) == %d\n", test, firstUniqChar(test))
	}
}

// Two approaches:
//	- Scan full string per char
//	- build char -> count map by scanning string once. Then scan string once left->right til first char found on map w/ 1 count.
func firstUniqChar(s string) int {
	charMap := map[rune]int{}

	for _, ch := range s {
		_, charCounted := charMap[ch]
		if !charCounted {
			charMap[ch] = 1
		} else {
			charMap[ch]++
		}
	}

	chIndex := -1
	for chIndex, ch := range s {
		chCount, ok := charMap[ch]
		if !ok {
			log.Fatalf("firstUniqChar(): char %q found in string %q but not in charMap %v\n", ch, s, charMap)
		} else {
			if chCount == 1 {
				return chIndex
			}
		}
	}

	// no non-repeating character found in string
	return chIndex
}
