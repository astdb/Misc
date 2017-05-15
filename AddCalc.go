
// Using only the add operator, implement multiply, subtract and division operations for integers. Results will also be ints.

package main

import (
    "fmt"
)

func main(){
    // test code
}

// a * b
func mul(a, b int) int {
	neg := false	// flag indicating if the end result should be negated
	if a < 0 {
		a = negate(a)
		neg = true
	}

	if b < 0 {
		b = negate(b)

		if neg {
			// if both operands are negative, don't need to negate the result
			neg = false
		} else {
			neg = true
		}
	}

	res := 0
	for i := 0; i < b; i++ {
		res = res + a
	}

	if neg {
		return negate(res)
	}

	return res
}

// a - b
func sub(a, b int) int {
	if b < 0 {
		return a + b
	}

	i := 1	
	for b < a {
		i++
	}

	return i
}

// a / b
func div(a, b int) int {
	// if b == 0, return an error

	if a < b {
		return 0
	}

	neg := false	// flag indicating if the end result should be negated
	if a < 0 {
		a = negate(a)
		neg = true
	}

	if b < 0 {
		b = negate(b)

		if neg {
			// if both operands are negative, don't need to negate the result
			neg = false
		} else {
			neg = true
		}
	}

	res := b
	i := 0
	for res < a {
		res = res + b
		i++
	}

	if neg {
		return negate(res)
	}

	return res
}

func negate(n int) int {
	neg := 0
	newSign := 0
	if n < 0 {
		newSign = 1
	} else {
		newSign = (-1)
	}

	for n != 0 {
		neg += newSign
		n += newSign
	}

	return neg
}

func abs(n int) {
	if n < 0 {
		return negate(n)
	}

	return n
}
