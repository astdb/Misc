/*
Given an array of integers arr, return true if and only if it is a valid mountain array.

Recall that arr is a mountain array if and only if:

arr.length >= 3
There exists some i with 0 < i < arr.length - 1 such that:
arr[0] < arr[1] < ... < arr[i - 1] < A[i]
arr[i] > arr[i + 1] > ... > arr[arr.length - 1]



Example 1:

Input: arr = [2,1]
Output: false
Example 2:

Input: arr = [3,5,5]
Output: false
Example 3:

Input: arr = [0,3,2,1]
Output: true


Constraints:

1 <= arr.length <= 104
0 <= arr[i] <= 104

*/

package main

import (
	"log"
)

func main() {
	tests := [][]int{{2, 1}, {3, 5, 5}, {0, 3, 2, 1}}

	for _, test := range tests {
		log.Printf("validMountainArray(%v) = %v\n", test, validMountainArray(test))
	}
}

func validMountainArray(arr []int) bool {
	peaked := false

	if len(arr) < 3 {
		return false
	}

	for i := 0; i < len(arr)-1; i++ {
		if !peaked {
			if arr[i+1] <= arr[i] {
				peaked = true
			}

		} else {
			if arr[i+1] >= arr[i] {
				return false
			}
		}
	}

	return true
}
