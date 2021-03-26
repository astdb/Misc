/*
Given an image represented by a NxN matrix, rotate it by 90degrees.

Can this be done in place?
*/

package main

import (
	"log"
)

func main() {
	// tests := [][][]int{{{1}}, {{1, 2}, {3, 4}}, {{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, {{1,2,3,4}, {5,6,7,8}, {9,10,11,12}, {13,14,15,16}}}
  tests := [][]int{{}, {1}, {1,2}, {1,2,3}}

	for _, test := range tests {
		// log.Printf("cols(%v) = %v\n", test, cols(test))
		// log.Printf("cols(%v) = %v\n", test, rotate(test))

    log.Printf("%v\n", revArr(test))
	}
}

func revArr(arr []int) []int {
  for i := 0; i < len(arr)/2;  i++ {
    arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]
  }

  return arr
}

func rotate(img [][]int) [][]int {
	imgCols := cols(img)

	for i := 0; i < len(img); i++ {
		// reverse img[i]
		for j := 0; j <= len(img[i])/2; j++ {
			imgCols[i][j], imgCols[i][len(imgCols[i])-1] = imgCols[i][len(imgCols[i])-1], imgCols[i][j]
		}
	}

	return imgCols
}

func cols(mat [][]int) [][]int {
	var res [][]int

	for i := 0; i < len(mat); i++ {
		res = append(res, []int{})
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			res[j] = append(res[j], mat[i][j])
		}
	}

	return res
}
