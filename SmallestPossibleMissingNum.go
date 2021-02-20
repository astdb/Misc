package main

import (
	"log"
	"sort"
)

func main() {
	tests := [][]int{{1, 3, 6, 4, 1, 2}, {1, 2, 3}, {-1, -3}}

	for _, test := range tests {
		log.Printf("smallestMissing(%v) = %d\n", test, smallestMissing(test))
	}
}

func smallestMissing(arr []int) int {
	sort.Ints(arr)

	j := 1
	for i := 0; i < len(arr); i++ {

		if arr[i] != j {
			// return j

		} else {
			if i < len(arr)-1 {
				if arr[i] == arr[i+1] {
					// next element should be compared to j unchanged
					// do nothing

				} else {
					// arr[i] == j, arr[i+1] must be compared to j+1
					j++

				}
			} else {
				// arr[i] == j, and its the last element in arr - j+1 must be the answer
				j++

			}
		}
	}

	return j
}
