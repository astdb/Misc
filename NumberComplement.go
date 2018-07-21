/*
Given a positive integer, output its complement number. The complement strategy is to flip the bits of its binary representation.

Note:
The given integer is guaranteed to fit within the range of a 32-bit signed integer.
You could assume no leading zero bit in the integer’s binary representation.
Example 1:
Input: 5
Output: 2
Explanation: The binary representation of 5 is 101 (no leading zero bits), and its complement is 010. So you need to output 2.
Example 2:
Input: 1
Output: 0
Explanation: The binary representation of 1 is 1 (no leading zero bits), and its complement is 0. So you need to output 0.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findComplement(5))
	fmt.Println(findComplement(1))
}

func findComplement(num int) int {
	// get binary form of num and store in slice (back to front, bits flipped)
	rem := 0
	num_binary := []int{}

	for num > 0 {
		rem = num % 2
		num = num / 2

		// flip bits
		if rem == 0 {
			rem = 1
		} else {
			rem = 0
		}

		num_binary = append(num_binary, rem)
	}

	// get decimal value of num_binary flipped
	decimalVal := 0
	for i := 0; i < len(num_binary); i++ {
		decimalVal += num_binary[i] * int(math.Pow(2.0, float64(i)))
	}

	return decimalVal
}
