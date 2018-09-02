// given two positive integers compute quotient (x/y) using only and addition/subtract/shift operators

package main

import (
	"fmt"
)

func main() {
	tests := [][]int{{4, 2}, {2, 4}, {4, 4}, {4, 0}, {0, 4}, {0, 0}, {9, 2}, {23545687, 565489}, {-1, -2}, {-5, -5}}
	for _, test := range tests {
		if test[1] > 0 {
			fmt.Println(test[0], "/", test[1], " == ", divide(test[0], test[1]))
		}
	}

}

// EPIJ solution
func divide(x, y int) int {
	result := 0
	var power uint32 = 32
	yPower := y << power

	for x >= y {
		for yPower > x {
			// yPower >>>= 1
			yPower >>= 1
			power--
		}

		result += 1 << power
		x -= yPower
	}

	return result
}

func quot(x, y int) int {
	res := 0

	if x == 0 && y == 0 {
		panic("0/0: undefined.")
	}

	if y == 0 {
		panic("Division by zero.")
	}

	if x < y {
		return 0
	}

	sign := -1
	if (x < 0 && y < 0) || (x > 0 && y > 0) {
		sign = 1
	}

	for x > 0 {
		x -= y
		if x >= 0 {
			res++
		}
	}

	// CANNOT USE *
	return res * sign
}
