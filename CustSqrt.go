
/* 
	Implement int sqrt(int x).
	Compute and return the square root of x.
	x is guaranteed to be a non-negative integer.
*/

package main

import (
    "fmt"
)

func main() {
	tests := []int{1,2,3,4,5,6,7,8,9}

	for _, testcase := range tests {
		fmt.Printf("%d\t%d\n", testcase, CustSqrt(testcase))
	}
}

func CustSqrt(x int) int {
	i := 0
	temp := 0

    for {
		sq := i * i

		if sq == x {
			return i
		}

		if sq > x {
			return temp
		}

		temp = i
		i++
	}
}
