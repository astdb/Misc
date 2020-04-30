/*
Write a program to check whether a given number is an ugly number.

Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.

Example 1:

Input: 6
Output: true
Explanation: 6 = 2 × 3
Example 2:

Input: 8
Output: true
Explanation: 8 = 2 × 2 × 2
Example 3:

Input: 14
Output: false
Explanation: 14 is not ugly since it includes another prime factor 7.
Note:

1 is typically treated as an ugly number.
Input is within the 32-bit signed integer range: [−231,  231 − 1].
*/

package main

import (
	"log"
)

func main() {
	tests := []int{0, 1, 6, 8, 14}

	for _, test := range tests {
		log.Printf("isUgly(%d) == %v\n", test, isUgly(test))
	}
}

func isUgly(num int) bool {
	loop := 0

	if num <= 0 {
		return false
	}

	if num == 1 {
		return true
	}

	for {
		if num%2 == 0 {
			num = num / 2
		} else if num%3 == 0 {
			num = num / 3
		} else if num%5 == 0 {
			num = num / 5
		} else if loop > 0 && num == 1 {
			return true
		} else {
			return false
		}

		loop++
	}
}
