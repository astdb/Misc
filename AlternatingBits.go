/*
Given a positive integer, check whether it has alternating bits: namely, if two adjacent bits will always have different values.

Example 1:
Input: 5
Output: True
Explanation:
The binary representation of 5 is: 101

Example 2:
Input: 7
Output: False
Explanation:
The binary representation of 7 is: 111.

Example 3:
Input: 11
Output: False
Explanation:
The binary representation of 11 is: 1011.

Example 4:
Input: 10
Output: True
Explanation:
The binary representation of 10 is: 1010.

*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(hasAlternatingBits(5))
	fmt.Println(hasAlternatingBits(7))
	fmt.Println(hasAlternatingBits(11))
	fmt.Println(hasAlternatingBits(10))
}

func hasAlternatingBits(n int) bool {
	bits := []int{}

	var rem int

	for n > 0 {
		rem = n % 2
		n = n / 2
		bits = append(bits, rem)
	}

	if len(bits) <= 1 {
		// return false	// no alternating patterin in a single bit
		return true // by definition
	}

	for i := 0; i < len(bits); i++ {
		if i < len(bits)-1 {
			if bits[i] == bits[i+1] {
				return false
			}
		}
	}

	return true
}
