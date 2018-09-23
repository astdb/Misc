// given an array representing a nonnegative decimal integer D, return an array representing D+1.
// e.g. []int{1,2,9} => []int{1,3,0}

package main

import (
	"fmt"
)

func main() {
	tests := [][]int{{0}, {1}, {1, 0}, {1, 2, 9}, {9, 9, 9, 9}}
	for _, test := range tests {
		testCopy := make([]int, len(test))
		copy(testCopy, test)
		fmt.Println(testCopy, "+ 1 = ", Increment(test))
	}
}

func Increment(num []int) []int {
	carry := 0
	for i := len(num) - 1; i >= 0; i-- {
		if i == len(num)-1 {
			res := num[i] + 1
			if res < 10 {
				num[i] = res
				// return num
			} else {
				num[i] = res % 10
				carry = res / 10
			}
		} else {
			res := num[i] + carry
			if res < 10 {
				num[i] = res
				carry = 0
				// return num
			} else {
				num[i] = res % 10
				carry = res / 10
			}
		}
	}

	if carry == 0 {
		return num
	} else {
		return append([]int{carry}, num...)
	}

	// return num
}
