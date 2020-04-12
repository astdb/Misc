/*
Given a string S, we can transform every letter individually to be lowercase or uppercase to create another string.  Return a list of all possible strings we could create.

Examples:
Input: S = "a1b2"
Output: ["a1b2", "a1B2", "A1b2", "A1B2"]

Input: S = "3z4"
Output: ["3z4", "3Z4"]

Input: S = "12345"
Output: ["12345"]
Note:

S will be a string with length between 1 and 12.
S will consist only of letters or digits.
*/

package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	tests := []string{"abc", "1abc", "a1b2", "3z4", "12345"}

	for testNo, test := range tests {
		log.Printf("Test #%d: letterCasePermutation(%s) == %v\n", testNo, test, letterCasePermutation(test))
	}
}

// - Start with a result string slice containing the input string.
// - Iterate through the string character-by-character.
// - If the char is non-numeric, construct two versions of the string with that char in upper an lower cases, for every string that already exist in the results slice. Add those to the results slice if not already in it.
// - Return results slice.

func letterCasePermutation(str string) []string {
	res := []string{str}

	// firstCh := true
	for chIndex, ch := range str {
		if _, err := strconv.Atoi(string(ch)); err == nil {
			// ch not numeric - goto next char

		} else {
			tmpRes := res
			var candidateStr string

			if len(tmpRes) <= 0 {
				// var candidateStr string

				// uppercase
				candidateStr = fmt.Sprintf("%s%s%s", str[:chIndex], strings.ToUpper(string(ch)), str[chIndex+1:])

				if !strFound(candidateStr, res) {
					res = append(res, candidateStr)
				}

				// lowercase
				candidateStr = fmt.Sprintf("%s%s%s", str[:chIndex], strings.ToLower(string(ch)), str[chIndex+1:])

				if !strFound(candidateStr, res) {
					res = append(res, candidateStr)
				}
			} else {
				for _, str2 := range tmpRes {
					// var candidateStr string

					// uppercase
					candidateStr = fmt.Sprintf("%s%s%s", str2[:chIndex], strings.ToUpper(string(ch)), str2[chIndex+1:])

					if !strFound(candidateStr, res) {
						res = append(res, candidateStr)
					}

					// lowercase
					candidateStr = fmt.Sprintf("%s%s%s", str2[:chIndex], strings.ToLower(string(ch)), str2[chIndex+1:])

					if !strFound(candidateStr, res) {
						res = append(res, candidateStr)
					}
				}
			}
		}
	}

	return res
}

// Helper function to check if a constructed permutation candidate already exists in the results slice.
func strFound(needle string, hayStack []string) bool {
	for _, str := range hayStack {
		if str == needle {
			return true
		}
	}

	return false
}
