/*
Write a program to find the n-th ugly number.

Ugly numbers are positive numbers whose prime factors only include 2, 3, 5. 

Example:

Input: n = 10
Output: 12
Explanation: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 is the sequence of the first 10 ugly numbers.
Note:  

1 is typically treated as an ugly number.
n does not exceed 1690.
*/

package main

import (
	"log"
)

func main() {
	tests := []int{10, 1690}

	for _, test := range tests {
		log.Printf("nthUglyNumber(%d) == %v\n", test, nthUglyNumber(test))
	}
}

func nthUglyNumber(n int) int {
	unCount := 0
	i := 0
    for {
		if isUgly(i) {
			unCount++
			if unCount == n {
				return i
			}
		}
		i++
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
