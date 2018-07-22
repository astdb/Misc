/*
A matrix is Toeplitz if every diagonal from top-left to bottom-right has the same element.

Now given an M x N matrix, return True if and only if the matrix is Toeplitz.


Example 1:

Input:
matrix = [
  [1,2,3,4],
  [5,1,2,3],
  [9,5,1,2]
]
Output: True
Explanation:
In the above grid, the diagonals are:
"[9]", "[5, 5]", "[1, 1, 1]", "[2, 2, 2]", "[3, 3]", "[4]".
In each diagonal all elements are the same, so the answer is True.
Example 2:

Input:
matrix = [
  [1,2],
  [2,2]
]
Output: False
Explanation:
The diagonal "[1, 2]" has different elements.

Note:

matrix will be a 2D array of integers.
matrix will have a number of rows and columns in range [1, 20].
matrix[i][j] will be integers in range [0, 99].

Follow up:

What if the matrix is stored on disk, and the memory is limited such that you can only load at most one row of the matrix into the memory at once?
What if the matrix is so large that you can only load up a partial row into the memory at once?

*/

package main

import (
	"fmt"
)

func main() {
	i := [][]int{
		{1, 2, 3, 4},
		{5, 1, 2, 3},
		{9, 5, 1, 2},
	}

	j := [][]int{
		{1, 2},
		{2, 2},
	}

	k := [][]int{{1, 2}}

	l := [][]int{{1}, {2}}

	fmt.Println(isToeplitzMatrix(i))
	fmt.Println(isToeplitzMatrix(j))
	fmt.Println(isToeplitzMatrix(k))
	fmt.Println(isToeplitzMatrix(l))
}

func isToeplitzMatrix(matrix [][]int) bool {
	for i := 0; i < len(matrix); i++ {
		if i > 0 {
			for j := 1; j < len(matrix[i]); j++ {
				if matrix[i][j] != matrix[i-1][j-1] {
					return false
				}
			}
		}
	}

	return true
}
