/*
We have a list of points on the plane.  Find the K closest points to the origin (0, 0).

(Here, the distance between two points on a plane is the Euclidean distance.)

You may return the answer in any order.  The answer is guaranteed to be unique (except for the order that it is in.)



Example 1:

Input: points = [[1,3],[-2,2]], K = 1
Output: [[-2,2]]
Explanation:
The distance between (1, 3) and the origin is sqrt(10).
The distance between (-2, 2) and the origin is sqrt(8).
Since sqrt(8) < sqrt(10), (-2, 2) is closer to the origin.
We only want the closest K = 1 points from the origin, so the answer is just [[-2,2]].
Example 2:

Input: points = [[3,3],[5,-1],[-2,4]], K = 2
Output: [[3,3],[-2,4]]
(The answer [[-2,4],[3,3]] would also be accepted.)


Note:

1 <= K <= points.length <= 10000
-10000 < points[i][0] < 10000
-10000 < points[i][1] < 10000

*/

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	tests := [][][]int{{{1, 3}, {-2, 2}}, {{3, 3}, {5, -1}, {-2, 4}}}
	i := 1
	for _, test := range tests {
		fmt.Println(kClosest(test, i))
		i++
	}
}

func kClosest(points [][]int, K int) [][]int {
	// initialize and populate list of points with distance to center calculated
	pvd := []*PointWithDist{}

	for _, point := range points {
		pvd = append(pvd, NewDistPoint(point, GetCenterDist(point)))
	}

	// sort slice of points with distances by distance
	sort.Slice(pvd, func(i, j int) bool {
		return pvd[i].Dist < pvd[j].Dist
	})

	// build slice of K closest points and return result
	result := [][]int{}
	for i := 0; i < K && i < len(pvd); i++ {
		result = append(result, pvd[i].Coords)
	}

	return result
}

// helper function to get distance to the center from a given cartesian point
func GetCenterDist(point []int) float64 {
	x := point[0]
	y := point[1]

	return math.Sqrt((float64(x*x) + float64(y*y)))
}

// struct representing cartesian point with distance to center
type PointWithDist struct {
	Coords []int
	Dist   float64
}

// 'constructor' for a point with distance
func NewDistPoint(coords []int, dist float64) *PointWithDist {
	var DP PointWithDist
	DP.Coords = coords
	DP.Dist = dist
	return &DP
}
