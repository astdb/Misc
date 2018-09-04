/*
An array is monotonic if it is either monotone increasing or monotone decreasing.
An array A is monotone increasing if for all i <= j, A[i] <= A[j].  An array A is monotone decreasing if for all i <= j, A[i] >= A[j].
Return true if and only if the given array A is monotonic.
*/

package main

import (
	"fmt"
)

func main() {
	tests := [][]int{nil, {}, {1}, {1, 2, 2, 3}, {6, 5, 4, 4}, {1, 3, 2}, {1, 2, 4, 5}, {1, 1, 1}}

	for _, test := range tests {
		fmt.Println(test, monot(test))
	}

}

func monot(x []int) bool {
	if len(x) == 0 {
		return true
	}

	monot_inc := false
	monot_dec := false
	for i := 1; i < len(x); i++ {
		if x[i] > x[i-1] {
			if monot_dec == true {
				// increasing element found in previously decreasing array
				return false
			}

			monot_inc = true
		}

		if x[i] < x[i-1] {
			if monot_inc == true {
				// decreasing element found in previously increasing array
				return false
			}

			monot_dec = true
		}
	}

	if monot_inc == false && monot_dec == false {
		// array consisted of same value
		return true
	}

	return monot_inc || monot_dec
}
