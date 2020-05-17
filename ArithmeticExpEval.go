/*
Given a string containing a list of numbers separated by arithmetic operators, write a program to output the correct result.

Example 1:
	Input: "1 + 2"
	Output: 3

Example 2:
	Input: "3 + 2 - 1 * 6 / 5 - 1 + 8"
	Output: 10.8

Notes:
	You may assume operators and operands are separated by one or more spaces.
	The only operators would be +,-,*,/ and the operands would be decimal digits.
*/

package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	tests := []string{"1 + 2", "1 * 2", "2 / 3", "22 / 7", "3 + 2 - 1", "4 / 2", "4 / 2 + 1", "4 / 2 - 1", "3 + 2 - 1 * 6 / 5 - 1 + 8"}
	for _, test := range tests {
		log.Printf("evaluate(%s) == %f\n", test, evaluate(test))
	}
}

func evaluate(expr string) float64 {
	// exprRunes = []rune(expr)
	// for i := 0; i < len(exprRunes); i++ {
	// 	if expRunes
	// }

	tokens := strings.Split(expr, " ")

	for _, token := range tokens {
		token = strings.TrimSpace(token)

		if token == "*" {
			// https://golang.org/pkg/strings/#SplitAfterN
			// fmt.Printf("%q\n", strings.SplitAfterN("a,b,c,d,e,f", ",", 2))
			// == ["a," "b,c,d,e,f"]
			comps := strings.SplitAfterN(expr, "*", 2)
			firstSubExp := comps[0]
			secondSubExp := comps[1]

			// firstSubExpRunes will be something like "subexp *" - get rid of that trailing *
			firstSubExpRunes := []rune(firstSubExp)
			firstSubExpRunes = firstSubExpRunes[:len(firstSubExpRunes)-1]
			firstSubExp = string(firstSubExpRunes)

			return evaluate(firstSubExp) * evaluate(secondSubExp)
		}

		if token == "/" {
			comps := strings.SplitAfterN(expr, "/", 2)
			firstSubExp := comps[0]
			secondSubExp := comps[1]

			// firstSubExpRunes will be something like "subexp *" - get rid of that trailing *
			firstSubExpRunes := []rune(firstSubExp)
			firstSubExpRunes = firstSubExpRunes[:len(firstSubExpRunes)-1]
			firstSubExp = string(firstSubExpRunes)

			return evaluate(firstSubExp) / evaluate(secondSubExp)
		}
	}

	// no *'s or /'s found
	for _, token := range tokens {
		if token == "+" {
			// https://golang.org/pkg/strings/#SplitAfterN
			// fmt.Printf("%q\n", strings.SplitAfterN("a,b,c,d,e,f", ",", 2))
			// == ["a," "b,c,d,e,f"]
			comps := strings.SplitAfterN(expr, "+", 2)
			firstSubExp := comps[0]
			secondSubExp := comps[1]

			// firstSubExpRunes will be something like "subexp *" - get rid of that trailing *
			firstSubExpRunes := []rune(firstSubExp)
			firstSubExpRunes = firstSubExpRunes[:len(firstSubExpRunes)-1]
			firstSubExp = string(firstSubExpRunes)

			return evaluate(firstSubExp) + evaluate(secondSubExp)
		}

		if token == "-" {
			// https://golang.org/pkg/strings/#SplitAfterN
			// fmt.Printf("%q\n", strings.SplitAfterN("a,b,c,d,e,f", ",", 2))
			// == ["a," "b,c,d,e,f"]
			comps := strings.SplitAfterN(expr, "-", 2)
			firstSubExp := comps[0]
			secondSubExp := comps[1]

			// firstSubExpRunes will be something like "subexp *" - get rid of that trailing *
			firstSubExpRunes := []rune(firstSubExp)
			firstSubExpRunes = firstSubExpRunes[:len(firstSubExpRunes)-1]
			firstSubExp = string(firstSubExpRunes)

			return evaluate(firstSubExp) - evaluate(secondSubExp)
		}
	}

	// no operators (/ * + -) found
	exprVal, err := strconv.ParseFloat(strings.TrimSpace(expr), 64)
	if err != nil {
		log.Fatal(err)
	}

	return exprVal
}
