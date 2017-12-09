
/*
	Determine whether an integer is a palindrome. Do this without extra space.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	tests := []int{0, 2, 3453, 3456543, 300, 5005005, 11}

	fmt.Println(math.MaxInt32)
	fmt.Println(math.MaxInt64)
	
	for _, v := range tests {
		fmt.Printf("%d -> %v\n", v, isPalindromeConstSpace(v))
	}
}

func isPalindromeConstSpace(x int) bool {
	// find divisor
	divisor := 1
	for x / divisor >= 10 {
		divisor *= 10
	}

	for x != 0 {
		if x < 0 {
			x = x * (-1)
		}
		
		// extract leading digit
		leading := x / divisor

		// extract trailing digit
		trailing := x % 10

		if leading == trailing {
			// there's hope

			// trim leading and trailing digits
			x = (x % divisor) / 10

			// reduce divisor appropriately as x being trimmed by two digits
			divisor = divisor / 100

			continue
		}

		// leading != trailing-  x cannnot be a 'palindrome' number
		return false
	}

	return true
}

func isPalindrome(x int) bool {
	if x < 0 {
		x = x * (-1)
	}

	// turn int into slice of digits
	xrem := x % 10
	x = x / 10

	xdigits := []int{xrem}

	for x > 0 {
		xrem = x % 10
		x = x / 10
	
		xdigits = append(xdigits, xrem)
	}

	i := 0
	j := len(xdigits)
	
	for j > i {
		if xdigits[i] == xdigits[j-1] {
			i++
			j--
			continue
		}

		return false
	}

	return true
}
