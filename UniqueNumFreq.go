/*
Given an array of integers arr, write a function that returns true if and only if the number of occurrences of each value in the array is unique.



Example 1:

Input: arr = [1,2,2,1,1,3]
Output: true
Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values have the same number of occurrences.
Example 2:

Input: arr = [1,2]
Output: false
Example 3:

Input: arr = [-3,0,1,-3,1,1,1,-3,10,0]
Output: true

*/

package main

import (
	"log"
)

func main() {
	tests := [][]int{{1, 2, 2, 1, 1, 3}, {1, 2}, {-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, {1, 2, 3, 4, 5, 6, 7, 8, 9}, {1, 2, 2, 1, 1, 3}}

	for _, test := range tests {
		log.Printf("uniqueOccurrences(%v) = %v\n", test, uniqueOccurrences(test))
	}
}

// 01. Scan full array for second occurrences of each element O(n^2) time / O(1) space
// 02. Sort array and check for similar adjacent values
// 03. Insert each value into a hashmap and detect
func uniqueOccurrences(arr []int) bool {
	// create map holding frequency of each arr elem
	valMap := map[int]int{}

	for _, v := range arr {
		_, seen := valMap[v]
		if seen {
			valMap[v]++
		} else {
			valMap[v] = 1
		}
	}

	// build map with element frequency as key - allowing detection of two similar frequencies
	occMap := map[int]bool{}

	for _, occ := range valMap {
		_, seen := occMap[occ]

		if seen {
			// more than one value in arr occur in the same number of times
			return false
		} else {
			occMap[occ] = true
		}
	}

	return true
}
