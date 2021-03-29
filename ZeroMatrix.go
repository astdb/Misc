/*
Given an NxM matrix, set each column and row to zero if any element within is zero.

*/

package main

import (
	"log"
)

func main() {
	tests := [][][]int{{{1, 2, 3, 4}, {5, 6, 0, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}, {{1,2,3,4,5},{6,7,0,9,10},{11,12,13,14,15}}}

	for _, test := range tests {
		log.Println("===================================================")
		printMatrix(test)
		log.Println("-------------")
		// log.Printf("%v\n", zeroMatrix(test))
		zeroedMat := zeroMatrix(test)
		printMatrix(zeroedMat)
	}
}

func zeroMatrix(mat [][]int) [][]int {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			// if an element is zero, set the start of that row and top of that column to zero
			if mat[i][j] == 0 {
				mat[i][0] = 0
				mat[0][j] = 0
			}
		}
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][0] == 0 || mat[0][j] == 0 {
				mat[i][j] = 0
			}
		}
	}

	return mat
}

func printMatrix(mat [][]int) {
	// log.Println(mat)

	for i := 0; i < len(mat); i++ {
		// for j := 0; j > len(mat[i]); j++ {
		//   log.Printf("%d ", mat[i][j])
		// }
		// log.Printf("\n")

		log.Println(mat[i])
	}
}
