// write a function to take an integer and returns an integer
// corresponding to the digits of the input written in reverse order
// e.g. 42 -> 24, -314 -> -413

package main

import (
	"fmt"
	"math"
)

func main() {
	tests := []int{24, -314}

	for _, test := range tests {
		fmt.Println("reverse(", test, ") == ", reverse(test))
		fmt.Println("rev(", test, ") == ", rev(test))
	}
}

// EPIJ
func rev(x int) int {
	res := 0
	rem := int(math.Abs(float64(x)))

	for rem != 0 {
		res = res*10 + rem%10
		rem /= 10
	}

	if x < 0 {
		return res * (-1)
	}

	return res
}

func reverse(x int) int {
	res := []int{}
	var rem int
	neg := false
	if x < 0 {
		x *= (-1)
		neg = true
	}

	for x > 0 {
		rem = x % 10
		x = x / 10
		res = append(res, rem)
	}

	x_rev := 0
	index := 0
	for i := len(res) - 1; i >= 0; i-- {
		x_rev += int(math.Pow10(i)) * res[index]
		index++
	}

	if neg == true {
		x_rev *= (-1)
	}

	return x_rev
}
