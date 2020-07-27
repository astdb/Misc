/*
Given a 2D integer matrix M representing the gray scale of an image, you need to design a smoother to make the gray scale of each cell becomes the average gray scale (rounding down) of all the 8 surrounding cells and itself. If a cell has less than 8 surrounding cells, then use as many as you can.

Example 1:
Input:
[[1,1,1],
 [1,0,1],
 [1,1,1]]
Output:
[[0, 0, 0],
 [0, 0, 0],
 [0, 0, 0]]
Explanation:
For the point (0,0), (0,2), (2,0), (2,2): floor(3/4) = floor(0.75) = 0
For the point (0,1), (1,0), (1,2), (2,1): floor(5/6) = floor(0.83333333) = 0
For the point (1,1): floor(8/9) = floor(0.88888889) = 0
Note:
The value in the given matrix is in the range of [0, 255].
The length and width of the given matrix are in the range of [1, 150].
*/

package main

import (
	"log"
)

func main() {
	tests := [][][]int{{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}}

	for _, test := range tests {
		log.Printf("imageSmoother(%v) == %v\n", test, imageSmoother(test))
	}
}

func imageSmoother(M [][]int) [][]int {
	result := [][]int{}

	// for each cell
	for i := 0; i < len(M); i++ {
		thisRow := []int{}
		log.Printf("row: %d\n", i)
		for j := 0; j < len(M[i]); j++ {
			log.Printf("\tcell: %d\n", j)

			surroundingCellCount := 0
			surroundingCellTotal := 0

			// self
			surroundingCellCount++
			surroundingCellTotal += M[i][j]

			// top left
			if ((i - 1) >= 0) && ((j - 1) >= 0) {
				log.Printf("\t\tTop Left\n")
				surroundingCellCount++
				surroundingCellTotal += M[i-1][j-1]
			}

			// top middle
			if ((i - 1) >= 0) && (len(M[i-1]) > j) {
				log.Printf("\t\tTop middle\n")

				surroundingCellCount++
				surroundingCellTotal += M[i-1][j]
			}

			// top right
			if ((i - 1) >= 0) && (len(M[i-1]) > (j + 1)) {

				log.Printf("\t\tTop right\n")
				surroundingCellCount++
				surroundingCellTotal += M[i-1][j+1]
			}

			// left
			if (j - 1) >= 0 {

				log.Printf("\t\tLeft\n")
				surroundingCellCount++
				surroundingCellTotal += M[i][j-1]
			}

			// right
			if (j + 1) < len(M[i]) {

				log.Printf("\t\tRight\n")
				surroundingCellCount++
				surroundingCellTotal += M[i][j+1]
			}

			// bottom left
			if (i+1) < len(M) && (j-1) >= 0 {

				log.Printf("\t\tBottom left\n")
				surroundingCellCount++
				surroundingCellTotal += M[i+1][j-1]
			}

			// bottom middle
			if (i+1) < len(M) && j < len(M[i+1]) {

				log.Printf("\t\tBottom mid\n")
				surroundingCellCount++
				surroundingCellTotal += M[i+1][j]
			}

			// bottom right
			if (i+1) < len(M) && (j+1) < len(M[i+1]) {

				log.Printf("\t\tBottom right\n")
				surroundingCellCount++
				surroundingCellTotal += M[i+1][j+1]
			}

			thisRow = append(thisRow, (surroundingCellTotal / surroundingCellCount))
		}

		result = append(result, thisRow)
	}

	// return surroundingCellTotal/surroundingCellCount
	return result
}
