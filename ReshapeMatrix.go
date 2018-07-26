/*
In MATLAB, there is a very useful function called 'reshape', which can reshape a matrix into a new one with different size but keep its original data.

You're given a matrix represented by a two-dimensional array, and two positive integers r and c representing the row number and column number of the wanted reshaped matrix, respectively.

The reshaped matrix need to be filled with all the elements of the original matrix in the same row-traversing order as they were.

If the 'reshape' operation with given parameters is possible and legal, output the new reshaped matrix; Otherwise, output the original matrix.

Example 1:
Input:
nums =
[[1,2],
 [3,4]]
r = 1, c = 4
Output:
[[1,2,3,4]]
Explanation:
The row-traversing of nums is [1,2,3,4]. The new reshaped matrix is a 1 * 4 matrix, fill it row by row by using the previous list.
Example 2:
Input:
nums =
[[1,2],
 [3,4]]
r = 2, c = 4
Output:
[[1,2],
 [3,4]]
Explanation:
There is no way to reshape a 2 * 2 matrix to a 2 * 4 matrix. So output the original matrix.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(matrixReshape(nil, 1, 4))
	fmt.Println(matrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4))
	fmt.Println(matrixReshape([][]int{{1, 2}, {3, 4}}, 2, 4))
	fmt.Println(matrixReshape([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}, 2, 8))
}

func matrixReshape(nums [][]int, r int, c int) [][]int {
	// check if reshape parameters are valid
	if len(nums) <= 0 || (r*c) != (len(nums)*len(nums[0])) {
		return nums
	}

	// expand original matrix
	expandedMatrix := []int{}

	for _, row := range nums {
		for _, element := range row {
			expandedMatrix = append(expandedMatrix, element)
		}
	}

	// fmt.Println(expandedMatrix)
	result := [][]int{}
	row := 0
	col := 0
	i := 0

	for row < r {
		result = append(result, []int{})
		for col < c {
			result[row] = append(result[row], expandedMatrix[i])
			col++
			i++
		}
		row++
		col = 0
	}

	return result
}
