/*
Given a string, output an integer score indicative of it's suitability for a password.
*/

package main

import (
	"fmt"
	"unicode"
)

func main() {
	tests := []string{"", "a", "Aa", "aaaaaaaaa", "Abcd1234", "P@$$w04d"}

	for _, v := range tests {
		fmt.Println("pwdComplexity(", v,") => ", pwdComplexity(v))
	}
}

func pwdComplexity(str string) int {
	str_runes := []rune(str)

	score := 0
	if len(str) < 8 {
		return score
	} else {
		score++
	}

	digits := false
	upper := false
	lower := false
	special := false
	for _, ch := range str_runes {
		if unicode.IsDigit(ch) {
			digits = true
		}

		if unicode.IsUpper(ch) {
			upper = true
		}

		if unicode.IsLower(ch) {
			lower = true
		}

		if unicode.IsSymbol(ch) || unicode.IsPunct(ch) {
			special = true
		}
	}

	if digits {
		score++
	}

	if upper && lower {
		score++
	}

	if special {
		score++
	}

	return score
}
