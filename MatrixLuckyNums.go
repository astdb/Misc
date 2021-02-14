/*
Given a m * n matrix of distinct numbers, return all lucky numbers in the matrix in any order.

A lucky number is an element of the matrix such that it is the minimum element in its row and maximum in its column.


Example 1:

Input: matrix = [[3,7,8],[9,11,13],[15,16,17]]
Output: [15]
Explanation: 15 is the only lucky number since it is the minimum in its row and the maximum in its column

Example 2:

Input: matrix = [[1,10,4,2],[9,3,8,7],[15,16,17,12]]
Output: [12]
Explanation: 12 is the only lucky number since it is the minimum in its row and the maximum in its column.

Example 3:

Input: matrix = [[7,8],[1,2]]
Output: [7]


Constraints:

m == mat.length
n == mat[i].length
1 <= n, m <= 50
1 <= matrix[i][j] <= 10^5.
All elements in the matrix are distinct.

*/

package main

import (
	"log"
)

func main() {
	tests := [][][]int{{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}, {{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}, {{7, 8}, {1, 2}}}

	for _, test := range tests {
		log.Printf("luckyNumbers(%v) = %v\n", test, luckyNumbers(test))
	}
}

func luckyNumbers(matrix [][]int) []int {
	// find min of each row
	mins := []int{}

	for _, row := range matrix {
		var thisMin int

		for i, v := range row {
			if i == 0 {
				thisMin = v
			} else if v < thisMin {
				thisMin = v
			}
		}

		mins = append(mins, thisMin)
	}

	log.Printf("Mins: %v\n", mins)

	// printout all cols
	maxs := []int{}
	for colIndex, _ := range matrix[0] {
		var thisMax int
		log.Printf("------------------------colindex: %d\n", colIndex)

		for i, row := range matrix {
			log.Printf("i = %d, row: %v\n", i, row)

			if i == 0 {
				log.Printf("Initializing thisMax to row[%d] (%d)..", colIndex, row[colIndex])
				thisMax = row[colIndex]
			} else if row[colIndex] > thisMax {
				log.Printf("New max found (row[%d] (%d)) - setting to thisMax", colIndex, row[colIndex])
				thisMax = row[colIndex]
			}

			// log.Printf("%d ", row[colIndex])
		}

		maxs = append(maxs, thisMax)

		// log.Printf("\n")
	}

	log.Printf("Maxs: %v\n", maxs)

	magic := []int{}
	for rowIndex, row := range matrix {
		for colIndex, elem := range row {
			if mins[rowIndex] == elem && maxs[colIndex] == elem {
				magic = append(magic, elem)
			}
		}
	}

	return magic

	/* // find max of each col
	maxes := []int{}
	for colIndex, _ := range matrix[0] {
	  var thisMax int
	  log.Printf("colIndex = %d\n", colIndex)

	  for i, v := range matrix {
	    log.Printf("i = %d\n", i)
	    if i == 0 {
	      log.Printf("Setting thisMax to v[%d] (%d)\n", colIndex, v[colIndex])
	      thisMax = v[colIndex]
	    }

	    if v[colIndex] > thisMax {
	      log.Printf("New max found")
	      thisMax = v[colIndex]
	    }
	  }

	  maxes = append(maxes, thisMax)
	}

	log.Printf("Maxes: %v\n", maxes)

	res := []int{}
	return res */
}
