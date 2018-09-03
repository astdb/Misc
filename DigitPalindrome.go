// check if an input integer's decimal form is a palindrome

package main

import (
	"fmt"
)

func main() {
	tests := []int{456, 454, 498894, 2, 0, 1,7,11,121,333,2147447412,-1,12,100,2147483647}

	for _, test := range tests {
		fmt.Println(test, isPal(test))
	}
}

func isPal(x int) bool {
	x_digits := []int{}
	var rem int

	if x < 0 {
		return false
	}

	for x > 0 {
		rem = x % 10
		x = x / 10
		x_digits = append(x_digits, rem)
	}

	startIndex := 0
	endIndex := len(x_digits)-1

	for startIndex <= endIndex {
		if x_digits[startIndex] != x_digits[endIndex] {
			return false
		}

		startIndex++
		endIndex--
	}

return true
}
