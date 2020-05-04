
/*
Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.


In Pascal's triangle, each number is the sum of the two numbers directly above it.

Example:

Input: 5
Output:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
*/

package main

import (
	"log"
)

func main() {
	tests := []int{2}

	for _, test := range tests {
		log.Printf("generate(%d) == %v\n", test, generate(test))
	}
}

func generate(numRows int) [][]int {
	result := [][]int{}
	
	for i := 0; i < numRows; i++ {
		// get the prev row
		var prevRow []int
		if i == 0 {
			thisRow := []int{1}
			result = append(result, thisRow)
		} else {
			prevRow = result[len(result)-1]
			// thisRow := []int{0,0,0,0}
			thisRow := []int{}

			// for i := 0; i < len(prevRow); i++ {
			// for j := 0; j < len(thisRow); j++ {
			for j := 0; j < (i+1); j++ {
				if j == 0 {
					// thisRow[j] = prevRow[j]
					thisRow = append(thisRow, prevRow[j])

				} else if j == len(thisRow)-1 {
					// thisRow[j] = prevRow[len(prevRow)-1]
					thisRow = append(thisRow, prevRow[len(prevRow)-1])

				} else {
					// thisRow[j] = prevRow[j] + prevRow[j-1]
					thisRow = append(thisRow, (prevRow[j] + prevRow[j-1]))
				}
			}

			result = append(result, thisRow)
		}
	}

	return result
}
