/*
Given a matrix A, return the transpose of A.

The transpose of a matrix is the matrix flipped over it's main diagonal, switching the row and column indices of the matrix.



Example 1:

Input: [[1,2,3],[4,5,6],[7,8,9]]
Output: [[1,4,7],[2,5,8],[3,6,9]]
Example 2:

Input: [[1,2,3],[4,5,6]]
Output: [[1,4],[2,5],[3,6]]

*/

package main

import (
	"fmt"
)

func main() {
	input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(input)
	fmt.Println(transpose(input))
}

func transpose(A [][]int) [][]int {
	if A == nil {
		return nil
	}

	x := len(A[0])
	T := [][]int{}

	for i := 0; i < x; i++ {
		b := []int{}
		T = append(T, b)
		
		for j := 0; j < len(A); j++ {
			T[i] = append(T[i], A[j][i])
		}
	}

	return T
}
