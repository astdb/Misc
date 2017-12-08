/*
	Given a 32-bit signed integer, reverse digits of an integer.

	Example 1:

	Input: 123
	Output:  321
	Example 2:

	Input: -123
	Output: -321
	Example 3:

	Input: 120
	Output: 21

	Note:
	Assume we are dealing with an environment which could only hold integers within the 32-bit signed integer range. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	testCases := []int{0, 00001, -1, 123, -123, 120, 1534236469}

	for _,v := range testCases {
		fmt.Printf("reverse(%d)\t -> %d\n", v, reverse(v))
	}
}

func reverse(x int) int {
	if x < 10 && x >= 0 {
		return x
	}

	xneg := false
	if x < 0 {
		xneg = true
		x *= (-1)
	}

	if x > math.MaxInt32 {
		return 0
	}

	xrem := x % 10
	x = x / 10

	xdigits := []int{xrem}

	for x > 0 {
		xrem = x % 10
		x = x / 10
	
		xdigits = append(xdigits, xrem)
	}

	rev := 0
	j := 0
	for i := len(xdigits)-1; i >= 0; i-- {
		rev += xdigits[j] * pow(10, i)
		j++
	}

	if xneg {
		return rev * (-1)
	}

	return rev
}

// simplified power function raising a positive int to 0 or positive int
func pow(x, y int) int {
	if y == 0 {
		return 1
	}

	pwr := x
	for i := 1; i < y; i++ {
		pwr *= x
	}

	return pwr
}
