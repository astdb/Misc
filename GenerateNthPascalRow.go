/*
Given a non-negative index k where k ≤ 33, return the kth index row of the Pascal's triangle.

Note that the row index starts from 0.


In Pascal's triangle, each number is the sum of the two numbers directly above it.

Example:

Input: 3
Output: [1,3,3,1]
Follow up:

Could you optimize your algorithm to use only O(k) extra space?

*/
package main

import (
	"log"
)

func main() {
	tests := []int{3}

	for _, test := range tests {
		log.Printf("getRow(%d) == %v\n", test, getRow(test))
	}
}

func getRow(rowIndex int) []int {
	var prevRow []int
	var thisRow []int

	for i := 0; i <= rowIndex; i++ {
		if i == 0 {
			thisRow = []int{1}

		} else {
			for j := 0; j < (i + 1); j++ {
				if j == 0 {
					thisRow = append(thisRow, prevRow[j])

				} else if j == i {
					thisRow = append(thisRow, prevRow[len(prevRow)-1])

				} else {
					thisRow = append(thisRow, (prevRow[j] + prevRow[j-1]))
				}
			}
		}

		if i == rowIndex {
			return thisRow
		} else {
			prevRow = thisRow
			thisRow = []int{}
		}
	}

	return thisRow
}
