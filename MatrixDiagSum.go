/*
Given a square matrix mat, return the sum of the matrix diagonals.

Only include the sum of all the elements on the primary diagonal and all the elements on the secondary diagonal that are not part of the primary diagonal.



Example 1:


Input: mat = [[1,2,3],
              [4,5,6],
              [7,8,9]]
Output: 25
Explanation: Diagonals sum: 1 + 5 + 9 + 3 + 7 = 25
Notice that element mat[1][1] = 5 is counted only once.
Example 2:

Input: mat = [[1,1,1,1],
              [1,1,1,1],
              [1,1,1,1],
              [1,1,1,1]]
Output: 8
Example 3:

Input: mat = [[5]]
Output: 5

*/

package main

import (
	"log"
)

func main() {
	tests := [][][]int{{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, {{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}, {{5}}}

	for _, test := range tests {
		log.Printf("diagonalSum(%v) == %d\n", test, diagonalSum(test))
	}
}

func diagonalSum(mat [][]int) int {
	// add primary diagonal
	// log.Printf("\tAdding primary diagonal..\n")
	diagSum := 0

	diagElemIdx := 0
	for _, row := range mat {
		if diagElemIdx < len(row) {
			// log.Printf("\t%dth row, %dth element (%d)", i, diagElemIdx, row[diagElemIdx])
			diagSum += row[diagElemIdx]
			diagElemIdx++
		}
	}

	// add secondary diagonal
	// log.Printf("\tAdding secondary diagonal..\n")
	if len(mat) > 0 && len(mat[0]) > 0 {
		diagElemIdx = len(mat[0]) - 1
	}

	// if the matrix size is odd, leave out the len(mat) / 2th element of secondary diagonal
	if len(mat)%2 == 0 {
		for _, row := range mat {
			if diagElemIdx >= 0 {
				diagSum += row[diagElemIdx]

			}

			diagElemIdx--
		}
	} else {
		for _, row := range mat {
			if diagElemIdx >= 0 && diagElemIdx != len(row)/2 {
				diagSum += row[diagElemIdx]

			}

			diagElemIdx--
		}
	}

	return diagSum
}
