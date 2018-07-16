/*
Let's call an array A a mountain if the following properties hold:

A.length >= 3
There exists some 0 < i < A.length - 1 such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]
Given an array that is definitely a mountain, return any i such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1].

Example 1:

Input: [0,1,0]
Output: 1
Example 2:

Input: [0,2,1,0]
Output: 1
*/

package main

import (
	"fmt"
)

func main() {
	tests := [][]int{{0,1,0},{0,2,1,0}}

	for _, test := range tests {
		fmt.Println(peakIndexInMountainArray(test))
	}
}

func peakIndexInMountainArray(A []int) int {
    if len(A) < 3 {
		return -1
	}

	for i := 0; i < len(A); i++ {
		if i+1 < len(A) {
			if A[i] > A[i + 1] {
				return i
			}
		}
	}

	return -1
}
