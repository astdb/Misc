/*
You are given a map in form of a two-dimensional integer grid where 1 represents land and 0 represents water. Grid cells are connected horizontally/vertically (not diagonally). The grid is completely surrounded by water, and there is exactly one island (i.e., one or more connected land cells). The island doesn't have "lakes" (water inside that isn't connected to the water around the island). One cell is a square with side length 1. The grid is rectangular, width and height don't exceed 100. Determine the perimeter of the island.

Example:

[[0,1,0,0],
 [1,1,1,0],
 [0,1,0,0],
 [1,1,0,0]]

Answer: 16
Explanation: The perimeter is the 16 yellow stripes in the image below:

*/

package main

import (
	"fmt"
)

func main() {
	testCase := [][]int{[]int{0, 1, 0, 0}, []int{1, 1, 1, 0}, []int{0, 1, 0, 0}, []int{1, 1, 0, 0}}
	fmt.Println(islandPerimeter(testCase))
}

func islandPerimeter(grid [][]int) int {
	// algorithm: for each '1' encountered, check if there are adjoining '0's, and add those up to form perimeter
	perimeter := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// considering grid[i][j]

			if grid[i][j] == 1 {
				// check west
				if j == 0 {
					perimeter++
				} else if grid[i][j-1] == 0 {
					perimeter++
				}

				// check north
				if i == 0 {
					perimeter++
				} else if grid[i-1][j] == 0 {
					perimeter++
				}

				// check east
				if j == len(grid[i])-1 {
					perimeter++
				} else if grid[i][j+1] == 0 {
					perimeter++
				}

				// check south
				if i == len(grid)-1 {
					perimeter++
				} else if grid[i+1][j] == 0 {
					perimeter++
				}
			}
		}
	}

	return perimeter
}
