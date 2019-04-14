/*
We are given two strings, A and B.
A shift on A consists of taking string A and moving the leftmost character to the rightmost position. For example, if A = 'abcde', then it will be 'bcdea' after one shift on A. Return True if and only if Acan become B after some number of shifts on A.
Example 1:
Input: A = 'abcde', B = 'cdeab'
Output: true

Example 2:
Input: A = 'abcde', B = 'abced'
Output: false

Note:
A and B will have length at most 100.

*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	tests := [][]string{{"abcde", "cdeab"}, {"abcde", "abced"}}

	for _, test := range tests {
		fmt.Printf("rotateString(%s, %s) == %v\n", test[0], test[1], rotateString(test[0], test[1]))
	}
}

func rotateString(A string, B string) bool {
	if A == B {
		return true
	}

	for x := len(A) - 1; x >= 0; x-- {
		if len(A) == len(B) && strings.Contains(B, A[x:]) && strings.Contains(B, A[:x]) {
			return true
		}
	}

	return false

}
