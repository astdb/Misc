/*
Given a string s and an integer array indices of the same length.

The string s will be shuffled such that the character at the ith position moves to indices[i] in the shuffled string.

Return the shuffled string.



Example 1:


Input: s = "codeleet", indices = [4,5,6,7,0,2,1,3]
Output: "leetcode"
Explanation: As shown, "codeleet" becomes "leetcode" after shuffling.
Example 2:

Input: s = "abc", indices = [0,1,2]
Output: "abc"
Explanation: After shuffling, each character remains in its position.
Example 3:

Input: s = "aiohn", indices = [3,1,4,2,0]
Output: "nihao"
Example 4:

Input: s = "aaiougrt", indices = [4,0,2,6,7,3,1,5]
Output: "arigatou"
Example 5:

Input: s = "art", indices = [1,0,2]
Output: "rat"


Constraints:

s.length == indices.length == n
1 <= n <= 100
s contains only lower-case English letters.
0 <= indices[i] < n
All values of indices are unique (i.e. indices is a permutation of the integers from 0 to n - 1).
*/

package main

import (
	"log"
)

func main() {
	tests := []*Blah{&Blah{S: "codeleet", Indices: []int{4, 5, 6, 7, 0, 2, 1, 3}}, &Blah{S: "abc", Indices: []int{0, 1, 2}}}

	for _, test := range tests {
		log.Printf("restoreString2(%s, %v) = %s\n", test.S, test.Indices, restoreString2(test.S, test.Indices))
	}
}

type Blah struct {
	S       string
	Indices []int
}

func restoreString2(s string, indices []int) string {
	s_runes := []rune(s)
	s_runes_copy := []rune(s)

	for i := 0; i < len(s_runes); i++ {
		if i < len(indices) && indices[i] < len(s_runes_copy) {
			s_runes_copy[indices[i]] = s_runes[i]
		}

	}

	return string(s_runes_copy)

}

func restoreString(s string, indices []int) string {
	resRunes := []rune{}

	for _, i := range indices {
		resRunes = append(resRunes, rune(i))
	}

	for i, ch := range s {
		resRunes[i] = ch
	}

	return string(resRunes)
}
