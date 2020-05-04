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
	tests := []int{1, 2, 3, 4, 5}

	for _, test := range tests {
		log.Printf("generate(%d) == %v\n", test, generate(test))
	}
}

func generate(numRows int) [][]int {
	result := [][]int{}

	for i := 0; i < numRows; i++ {
		// log.Printf("-----------------------------\ni == %d\n", i)
		// get the prev row
		var prevRow []int
		if i == 0 {
			// log.Println("Setting first row..")
			thisRow := []int{1}
			result = append(result, thisRow)
			// log.Printf("result == %v\n", result)
		} else {
			// log.Println("Setting non-first row")
			prevRow = result[len(result)-1]
			thisRow := []int{}

			for j := 0; j < (i + 1); j++ {
				// log.Printf("\t--------------\nj == %d\n", j)
				if j == 0 {
					// log.Println("\tSetting thisrow first element..")
					thisRow = append(thisRow, prevRow[j])
					// log.Printf("\tthisRow == %v\n", thisRow)

				} else if j == i {
					// log.Println("\tSetting thisrow last element..")
					thisRow = append(thisRow, prevRow[len(prevRow)-1])
					// log.Printf("\tthisRow == %v\n", thisRow)

				} else {
					// log.Println("\tSetting one of thisrow middle elements..")
					thisRow = append(thisRow, (prevRow[j] + prevRow[j-1]))
					// log.Printf("\tthisRow == %v\n", thisRow)
				}
			}

			result = append(result, thisRow)
		}
	}

	return result
}
