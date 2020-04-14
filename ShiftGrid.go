/*
Given a 2D grid of size m x n and an integer k. You need to shift the grid k times.

In one shift operation:

Element at grid[i][j] moves to grid[i][j + 1].
Element at grid[i][n - 1] moves to grid[i + 1][0].
Element at grid[m - 1][n - 1] moves to grid[0][0].
Return the 2D grid after applying shift operation k times.



Example 1:


Input: grid = [[1,2,3],[4,5,6],[7,8,9]], k = 1
Output: [[9,1,2],[3,4,5],[6,7,8]]
Example 2:


Input: grid = [[3,8,1,9],[19,7,2,5],[4,6,11,10],[12,0,21,13]], k = 4
Output: [[12,0,21,13],[3,8,1,9],[19,7,2,5],[4,6,11,10]]
Example 3:

Input: grid = [[1,2,3],[4,5,6],[7,8,9]], k = 9
Output: [[1,2,3],[4,5,6],[7,8,9]]


Constraints:

m == grid.length
n == grid[i].length
1 <= m <= 50
1 <= n <= 50
-1000 <= grid[i][j] <= 1000
0 <= k <= 100
*/

package main

import (
	"log"
)

func main() {
	var grid [][]int
	var k int

	grid = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	k = 1
	log.Printf("shiftGrid(%v, %d) == %v\n", grid, k, shiftGrid(grid, k))

	grid = [][]int{{3,8,1,9},{19,7,2,5},{4,6,11,10},{12,0,21,13}}
	k = 1
	log.Printf("shiftGrid(%v, %d) == %v\n", grid, k, shiftGrid(grid, k))
}

func shiftGrid(grid [][]int, k int) [][]int {
	if len(grid) <= 0 {
		log.Fatal("shiftGrid(): empty grid input.")
	}

	m := len(grid)
	n := len(grid[0])

	// declare and initialize target grid with dummy data
	grid2 := [][]int{}
	for i := 0; i < m; i++ {
		grid2 = append(grid2, []int{})
		for j := 0; j < n; j++ {
			grid2[i] = append(grid2[i], 0)
		}
	}

	// perform shift ops, copying shifted values to target grid from input grid
	for shiftCount := 0; shiftCount < k; shiftCount++ {
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				if (j + 1) < len(grid2[i]) {
					grid2[i][j+1] = grid[i][j]
				} else {
					if (i + 1) < len(grid2) {
						grid2[i+1][0] = grid[i][j]
					} else {
						grid2[0][0] = grid[i][j]
					}
				}
			}
		}
	}

	return grid2
}
