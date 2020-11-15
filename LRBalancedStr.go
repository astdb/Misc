/*
Balanced strings are those who have equal quantity of 'L' and 'R' characters.

Given a balanced string s split it in the maximum amount of balanced strings.

Return the maximum amount of splitted balanced strings.



Example 1:

Input: s = "RLRRLLRLRL"
Output: 4
Explanation: s can be split into "RL", "RRLL", "RL", "RL", each substring contains same number of 'L' and 'R'.
Example 2:

Input: s = "RLLLLRRRLR"
Output: 3
Explanation: s can be split into "RL", "LLLRRR", "LR", each substring contains same number of 'L' and 'R'.
Example 3:

Input: s = "LLLLRRRR"
Output: 1
Explanation: s can be split into "LLLLRRRR".
Example 4:

Input: s = "RLRRRLLRLL"
Output: 2
Explanation: s can be split into "RL", "RRRLLRLL", since each substring contains an equal number of 'L' and 'R'

*/

package main

import (
	"log"
)

func main() {
	tests := []string{"RLRRLLRLRL", "RLLLLRRRLR", "LLLLRRRR", "RLRRRLLRLL"}

	for _, test := range tests {
		log.Printf("balancedStringSplit(%s) = %d\n", test, balancedStringSplit(test))
	}
}

func balancedStringSplit(s string) int {
	// split string into slice of runes (unicode chars)
	sRunes := []rune(s)

	// iterate through s's runes, keping count of L's and R's seen so far
	// if the number of L's and R's seen are nonzero and equal, add one to the total of balanced strings.
	// if the given string is LR-balanced, start the output total at 1.
	totalBalanced := 0
	subStrFound := false

	if balanced(s) {
		totalBalanced++
	} else {
		return totalBalanced
	}

	LCount := 0
	RCount := 0
	for i := 0; i < len(sRunes); i++ {
		if sRunes[i] == 'L' {
			LCount++

			if LCount > 0 && RCount > 0 && LCount == RCount {
				if !subStrFound {
					subStrFound = true
				} else {
					totalBalanced++
				}
			}
		}

		if sRunes[i] == 'R' {
			RCount++

			if LCount > 0 && RCount > 0 && LCount == RCount {
				if !subStrFound {
					subStrFound = true
				} else {
					totalBalanced++
				}
			}
		}
	}

	return totalBalanced
}

func balanced(s string) bool {
	LCount := 0
	RCount := 0

	for _, ch := range s {
		if ch == 'L' {
			LCount++
		}

		if ch == 'R' {
			RCount++
		}
	}

	if LCount > 0 && RCount > 0 && LCount == RCount {
		return true
	}

	return false
}
