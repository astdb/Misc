// write a function to return the minimum number of parentheses required to add to a given string of parenthese to make it matched.
// e.g. ()() needs 0 additions
//      ()) needs 1

package main

import (
	"fmt"
)

func main() {
	TestStrings := []string{"(()())", "((())", "())", ""}

	for _, tc := range TestStrings {
		fmt.Printf("Test case: %s, match value: %d\n", tc, bracket_match(tc))
	}
}

func bracket_match(bracket_str string) int {
	count := 0

	for _, c := range bracket_str {
		open_r := '('
		close_r := ')'

		if c == open_r {
			count++
		} else if c == close_r {
			count--
		}
	}

	if count < 0 {
		return count * (-1)
	}

	return count
}

//!-
