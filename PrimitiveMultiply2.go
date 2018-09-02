// implement multiplication (x*y) without using arithmetic operators
// use only bitwise operators, assignment, equality checks, and boolean
// combinations thereof.

package main

import (
	"fmt"
)

func main() {
	// fmt.Println(multip(2,4))
	fmt.Println(multip2(3, 6))
}

// EPIJ solution
// initialize a result to zero and and iterate through the bits of x, adding (2^k)*y to the result
// if the kth bit of x is 1
func multip3(x, y int) int {
	sum := 0

	for x != 0 {
		// examine each bit of x
		if (x & 1) != 0 {
			sum = add(sum, y)
		}

		x >>= 1 // originally x >>>= 1
		y <<= 1
	}

	return sum
}

func add(a, b int) int {
	sum := 0
	carryin := 0
	k := 1
	tempA := a
	tempB := b

	for tempA != 0 || tempB != 0 {
		ak := a & k
		bk := b & k

		carryout := (ak & bk) | (ak & carryin) | (bk & carryin)
		sum |= (ak ^ bk ^ carryin)
		carryin = carryout << 1
		k <<= 1

		// both originally >>>=
		tempA >>= 1
		tempB >>= 1
	}

	return sum | carryin
}

// -----------------------------------------------

func multip2(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}

	if a == 1 {
		return b
	}

	if b == 1 {
		return a
	}

	tempA := a
	orReqd := false

	for b != 0 {
		if b == 1 {
			break
		}

		if (b & 1) == 1 {
			orReqd = true
		}

		a <<= 1 // double a
		b >>= 1 // halve b
	}

	if orReqd {
		return a | tempA
	}

	return a
}

// first attempt - USES ADDITION OP
func multip(x, y int) int {
	if y == 0 {
		return 0
	}

	//
	if y > 0 {
		return (x + multip(x, y-1))
	}

	if y < 0 {
		return -multip(x, -y)
	}

	return -1
}
