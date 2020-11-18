/*
A valid parentheses string is either empty (""), "(" + A + ")", or A + B, where A and B are valid parentheses strings, and + represents string concatenation.  For example, "", "()", "(())()", and "(()(()))" are all valid parentheses strings.

A valid parentheses string S is primitive if it is nonempty, and there does not exist a way to split it into S = A+B, with A and B nonempty valid parentheses strings.

Given a valid parentheses string S, consider its primitive decomposition: S = P_1 + P_2 + ... + P_k, where P_i are primitive valid parentheses strings.

Return S after removing the outermost parentheses of every primitive string in the primitive decomposition of S.



Example 1:

Input: "(()())(())"
Output: "()()()"
Explanation:
The input string is "(()())(())", with primitive decomposition "(()())" + "(())".
After removing outer parentheses of each part, this is "()()" + "()" = "()()()".
Example 2:

Input: "(()())(())(()(()))"
Output: "()()()()(())"
Explanation:
The input string is "(()())(())(()(()))", with primitive decomposition "(()())" + "(())" + "(()(()))".
After removing outer parentheses of each part, this is "()()" + "()" + "()(())" = "()()()()(())".
Example 3:

Input: "()()"
Output: ""
Explanation:
The input string is "()()", with primitive decomposition "()" + "()".
After removing outer parentheses of each part, this is "" + "" = "".

*/

package main

import (
	"log"
	"strings"
)

func main() {
	tests := [][]string{{"(()())(())", "()()()"}, {"(()())(())(()(()))", "()()()()(())"}, {"()()", ""}}

	for _, test := range tests {
		// log.Printf(" removeOuterParentheses(%s) = %s\n", test[0],  removeOuterParentheses(test))

		if removeOuterParentheses(test[0]) != test[1] {
			log.Printf("FAIL: expected %s for removeOuterParentheses(%s)\n", test[0], test[1])
		} else {
			log.Println("PASS")
		}
	}
}

func removeOuterParentheses(pStr string) string {
	var res strings.Builder

	// convert pStr into rune slice for per-character analysis
	pStrRunes := []rune(pStr)

	openCount := 0  // open parenths in current primitive
	closeCount := 0 // close parenths in current primitive
	startIndex := 0 // start index of current primitive
	closeIndex := 0 // close index of current primitive

	for i := 0; i < len(pStrRunes); i++ {
		ch := pStrRunes[i]

		if ch == '(' {
			openCount++
		}

		if ch == ')' {
			closeCount++
		}

		if openCount == closeCount {
			// end of primitive
			openCount = 0
			closeCount = 0

			closeIndex = i
			res.WriteString(pStr[startIndex+1 : closeIndex])

			startIndex = closeIndex + 1
		}
	}

	return res.String()
}
