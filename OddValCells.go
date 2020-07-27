/*
Given n and m which are the dimensions of a matrix initialized by zeros and given an array indices where indices[i] = [ri, ci]. For each pair of [ri, ci] you have to increment all cells in row ri and column ci by 1.

Return the number of cells with odd values in the matrix after applying the increment to all indices.

 

Example 1:


Input: n = 2, m = 3, indices = [[0,1],[1,1]]
Output: 6
Explanation: Initial matrix = [[0,0,0],[0,0,0]].
After applying first increment it becomes [[1,2,1],[0,1,0]].
The final matrix will be [[1,3,1],[1,3,1]] which contains 6 odd numbers.
Example 2:


Input: n = 2, m = 2, indices = [[1,1],[0,0]]
Output: 0
Explanation: Final matrix = [[2,2],[2,2]]. There is no odd number in the final matrix.
 

Constraints:

1 <= n <= 50
1 <= m <= 50
1 <= indices.length <= 100
0 <= indices[i][0] < n
0 <= indices[i][1] < m
*/

package main

import (
  "log"
)

func main() {
  log.Printf("oddCells(%d, %d, %v) == %d\n", 2, 3, [][]int{{0,1},{1,1}}, oddCells(2, 3, [][]int{{0,1},{1,1}}))
  log.Printf("oddCells(%d, %d, %v) == %d\n", 2, 2, [][]int{{1,1},{0,0}}, oddCells(2, 2, [][]int{{1,1},{0,0}}))
}

func oddCells(n int, m int, indices [][]int) int {
  // initialize matrix
  mat := [][]int{}
  for i := 0; i < n; i++ {
    row := []int{}
    for j := 0; j < m; j++ {
      row = append(row, 0)
    }

    mat = append(mat, row)
  }

  for _, index := range indices {
    if len(index) >= 2 && index[0] < n && index[1] < m {
      mat[index[0]][index[1]]++
    }
  }

  log.Printf("\toddCells(): %v\n", mat)

  res := 0
  for i := 0; i < n; i++ {
    for j := 0; j < m; j++ {
      if mat[i][j] % 2 == 1 {
        // odd element
        res++
      }
    }
  }

  return res
}
