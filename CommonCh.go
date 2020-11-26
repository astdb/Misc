/*
Given an array A of strings made only from lowercase letters, return a list of all characters that show up in all strings within the list (including duplicates).  For example, if a character occurs 3 times in all strings but not 4 times, you need to include that character three times in the final answer.

You may return the answer in any order.



Example 1:

Input: ["bella","label","roller"]
Output: ["e","l","l"]
Example 2:

Input: ["cool","lock","cook"]
Output: ["c","o"]


Note:

1 <= A.length <= 100
1 <= A[i].length <= 100
A[i][j] is a lowercase letter

*/

package main

import (
	"log"
)

func main() {
	tests := [][]string{{"bella", "label", "roller"}, {"cool", "lock", "cook"}}

	for _, test := range tests {
		log.Printf("commonChars(%v) = %v\n", test, commonChars(test))
	}
}

// create map of char to []int - keepint track of how may times a given char appears in each string
func commonChars(A []string) []string {
	charCountsG := map[rune][]int{}

	for _, str := range A {
		charCountsLoc := map[rune]int{}

		for _, ch := range str {
			_, counted := charCountsLoc[ch]
			if counted {
				charCountsLoc[ch]++
			} else {
				charCountsLoc[ch] = 1
			}
		}

		// update global char counts
		for ch, locCount := range charCountsLoc {
			_, charSeen := charCountsG[ch]
			if charSeen {
				charCountsG[ch] = append(charCountsG[ch], locCount)
			} else {
				charCountsG[ch] = []int{locCount}
			}
		}
	}

	result := []string{}

	for ch, counts := range charCountsG {
		if len(counts) > 0 && len(counts) == len(A) {
			for i := 0; i < getMin(counts); i++ {
				result = append(result, string(ch))
			}
		}
	}

	return result
}

func getMin(n []int) int {
	var min int

	for i := 0; i < len(n); i++ {
		if i == 0 {
			min = n[i]
		} else if n[i] < min {
			min = n[i]
		}
	}

	return min
}
