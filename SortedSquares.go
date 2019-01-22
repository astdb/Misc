/*
Given an array of integers A sorted in non-decreasing order, return an array of the squares of each number, also in sorted non-decreasing order.

 

Example 1:

Input: [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Example 2:

Input: [-7,-3,2,3,11]
Output: [4,9,9,49,121]
 

Note:

1 <= A.length <= 10000
-10000 <= A[i] <= 10000
A is sorted in non-decreasing order.

*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	tests := [][]int{{-4,-1,0,3,10}, {-7,-3,2,3,11}}

	for _, test := range tests {
		fmt.Println(sortedSquares(test))
	}
}

func sortedSquares(nums []int) []int {
	// build list of squared values
	squared := []int{}
	for _, n := range nums {
		squared = append(squared, n*n)
	}

	// return sorted squared values
	sort.Ints(squared)
	return squared
}
