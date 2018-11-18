/*
On a N * N grid, we place some 1 * 1 * 1 cubes.

Each value v = grid[i][j] represents a tower of v cubes placed on top of grid cell (i, j).

Return the total surface area of the resulting shapes.



Example 1:

Input: [[2]]
Output: 10
Example 2:

Input: [[1,2],[3,4]]
Output: 34
Example 3:

Input: [[1,0],[0,2]]
Output: 16
Example 4:

Input: [[1,1,1],[1,0,1],[1,1,1]]
Output: 32
Example 5:

Input: [[2,2,2],[2,1,2],[2,2,2]]
Output: 46


Note:

1 <= N <= 50
0 <= grid[i][j] <= 50
*/

package main

import (
	"fmt"
)

func main() {
	testCases := [][][]int{[][]int{[]int{2}}, [][]int{[]int{1, 2}, []int{3, 4}}, [][]int{[]int{1, 0}, []int{0, 2}}, [][]int{[]int{1, 1, 1}, []int{1, 0, 1}, []int{1, 1, 1}}, [][]int{[]int{2, 2, 2}, []int{2, 1, 2}, []int{2, 2, 2}}}

	for _, test := range testCases {
		fmt.Println(surfaceArea(test))
	}
}

func surfaceArea(grid [][]int) int {
	// algorithm: for each grid position, check the positions around it (north, east, south, and west).
	// if there's nothing on a side or it's a grid edge, add the height of the current position to total (area of that
	// side). If there's something on a side position and it's taller than current, no area is added to total.
	// If an adjacent position has a height smaller than the current position,
	// add the difference of heights for that side to the current total. This should be done for each of the four
	// sides. Also, for each nonzero grid point, add two area units to the total, representing the top and botton areas.
	totalArea := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > 0 {
				totalArea += 2 // mandatory top/bottom

				// check north face
				if i == 0 {
					// on the northern edge
					totalArea += grid[i][j]
				} else if grid[i-1][j] == 0 {
					// nothing on the north
					totalArea += grid[i][j]
				} else if grid[i-1][j] > 0 {
					// there's something on the north..
					if grid[i-1][j] < grid[i][j] {
						// ..and it's shorter
						totalArea += (grid[i][j] - grid[i-1][j])
					} // else if grid[i-1][j] > grid[i][j] {
					// 	totalArea += grid[i-1][j] - grid[i][j]
					// }
				}

				// check southern face
				if i == len(grid)-1 {
					// on the southern edge
					totalArea += grid[i][j]
				} else if grid[i+1][j] == 0 {
					// nothing on south side
					totalArea += grid[i][j]
				} else if grid[i+1][j] > 0 {
					// there's something on the south side..
					if grid[i+1][j] < grid[i][j] {
						// .. and it's shorter
						totalArea += (grid[i][j] - grid[i+1][j])
					} // else if grid[i+1][j] > grid[i][j] {
					// 	totalArea += grid[i+1][j] - grid[i][j]
					// }
				}

				// check eastern face
				if j == 0 {
					// on the eastern edge
					totalArea += grid[i][j]
				} else if grid[i][j-1] == 0 {
					// nothing on the eastern edge
					totalArea += grid[i][j]
				} else if grid[i][j-1] > 0 {
					// there's something on the eastern side..
					if grid[i][j-1] < grid[i][j] {
						// .. and it's shorter
						totalArea += (grid[i][j] - grid[i][j-1])
					}
				}

				// check western face
				if j == len(grid[i])-1 {
					// on the western edge
					totalArea += grid[i][j]
				} else if grid[i][j+1] == 0 {
					// nothing on the western side
					totalArea += grid[i][j]
				} else if grid[i][j+1] > 0 {
					// there's something on the western side..
					if grid[i][j+1] < grid[i][j] {
						// .. and it's shorter
						totalArea += grid[i][j] - grid[i][j+1]
					}
				}
			}
		}
	}

	return totalArea
}
