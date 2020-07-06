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
