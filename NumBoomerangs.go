/*
Given n points in the plane that are all pairwise distinct, a "boomerang" is a tuple of points (i, j, k) such that the distance between i and j equals the distance between i and k (the order of the tuple matters).

Find the number of boomerangs. You may assume that n will be at most 500 and coordinates of points are all in the range [-10000, 10000] (inclusive).

Example:

Input:
[[0,0],[1,0],[2,0]]

Output:
2

Explanation:
The two boomerangs are [[1,0],[0,0],[2,0]] and [[1,0],[2,0],[0,0]]
*/

package main

import (
  "log"
)

func main() {
  
}

// - generate all 3-point permutations possible from n points.
// - for each, calculate distances and determine if boomerang. 
// - store distance for each new pair calculated to eliminate recalculation effort. 
func numberOfBoomerangs(points [][]int) int {
  
}
