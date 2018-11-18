/*
Given a string S that only contains "I" (increase) or "D" (decrease), let N = S.length.

Return any permutation A of [0, 1, ..., N] such that for all i = 0, ..., N-1:

If S[i] == "I", then A[i] < A[i+1]
If S[i] == "D", then A[i] > A[i+1]


Example 1:

Input: "IDID"
Output: [0,4,1,3,2]
Example 2:

Input: "III"
Output: [0,1,2,3]
Example 3:

Input: "DDI"
Output: [3,2,0,1]


Note:

1 <= S.length <= 10000
S only contains characters "I" or "D".
*/

package main

import (
	"fmt"
)

func main() {
	testCases := []string{"IDID", "III", "DDI"}

	for _, test := range testCases {
		fmt.Println(diStringMatch(test))
	}
}

func diStringMatch(str string) []int {
	res := []int{}

	if len(str) > 0 {
		str_runes := []rune(str)

		if str_runes[0] == 'D' {
			res = append(res, len(str))
		} else if str_runes[0] == 'I' {
			res = append(res, 0)
		}

		for _, ch := range str {
			if ch == 'I' {
				res = append(res, res[len(res)-1]+1)
			}

			if ch == 'D' {
				res = append(res, res[len(res)-1]-1)
			}
		}
	}

	return res
}
