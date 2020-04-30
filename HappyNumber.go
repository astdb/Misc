/*
Write an algorithm to determine if a number n is "happy".

A happy number is a number defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.

Return True if n is a happy number, and False if not.

Example:

Input: 19
Output: true
Explanation:
1^2 + 9^2 = 82
8^2 + 2^2 = 68
6^2 + 8^2 = 100
1^2 + 0^2 + 0^2 = 1
*/

package main

import (
	"fmt"
	"log"
)

func main() {
	tests := []int{15, 16, 17, 18, 19, 20, 21, 22}

	for _, test := range tests {
		log.Printf("isHappy(%d) == %v\n", test, isHappy(test))
		// log.Printf("getDigits(%d) == %v\n", test, getDigits(test))
	}
}

func isHappy(n int) bool {
	nValuesFound := map[int]bool{}
	for {
		digSqTot := 0
		digits := getDigits(n)
		log.Printf("n = %d\n", n)
		for i, digit := range digits {
			digSqTot = digSqTot + sq(digit)
			// log.Printf("%d^2")
			if i > 0 {
				fmt.Printf(" + %d", digit)
			} else {
				fmt.Printf("%d", digit)
			}
		}
		fmt.Printf(" = %d\n", digSqTot)

		n = digSqTot

		if digSqTot == 1 {
			return true
		}

		_, found := nValuesFound[n]
		if found {
			return false
		} else {
			nValuesFound[n] = true
		}
	}

	return false
}

func sq(x int) int {
	return x * x
}

// return slice of decimal digits x consists of
func getDigits(x int) []int {
	res := []int{}

	rem := x % 10
	res = append(res, rem)
	x = x / 10

	for x > 0 {
		rem = x % 10
		res = append(res, rem)
		x = x / 10
	}

	return res
}
