package main

import (
	"log"
)

func main() {
	tests := [][]int32{{3, 2, 1, 2, 3}, {6}, {7, 1, 3, 4, 1, 7}}

	for testNo, test := range tests {
		log.Printf("Test #%d: minDists(%v) == %d\n", testNo, test, minimumDistances(test))
	}
}

// Complete the minimumDistances function below.

// Design: Keep a map of array values to their indices. Iterate through the array per element.
// For each value, check if the value is seen in the map - if not, insert it with its current index.
// If a value is seen in the map, deduct its recorded index from current element's index to calculate
// distance - if smaller than recorded minimum distance, update minimum distance varable, and remove value from map.
// Upon completing array iteration, return minimum distance value. 
func minimumDistances(a []int32) int32 {
	// map of similar values and their indexes
	vals := map[int32]int{}
	var minDist int32
	minDist = -1

	// var i int32
	for i := 0; i < len(a); i++ {
		// check if val seen before
		valIndex, seen := vals[a[i]]

		if seen {
			var thisDist int32
			thisDist = int32(i - valIndex)
			if minDist == -1 {
				minDist = thisDist
			} else if minDist > thisDist {
				minDist = thisDist
			}

			delete(vals, a[i])
		} else {
			vals[a[i]] = i
		}
	}

	return minDist
}
