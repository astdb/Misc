/*
Given a non-negative integer num, repeatedly add all its digits until the result has only one digit.

Example:

Input: 38
Output: 2
Explanation: The process is like: 3 + 8 = 11, 1 + 1 = 2.
             Since 2 has only one digit, return it.
Follow up:
Could you do it without any loop/recursion in O(1) runtime?
*/

package main

import (
	"fmt"
)

func main() {
	tests := []int{38}

	for _, test := range tests {
		fmt.Println(addDigits(test))
	}
}

func addDigits(num int) int {
	// if num has >1 decimal digits
	for num > 9 {
		// break num into decimal digits
		digits := []int{}
		rem := num % 10
		num = num / 10
		digits = append(digits, rem)
		for num > 0 {
			rem = num % 10
			num = num / 10
			digits = append(digits, rem)
		}

		// calculate total of num's decimal digits
		tot := 0
		for _, d := range digits {
			tot += d
		}

		// replace num with digit total
		num = tot
	}

	return num
}
