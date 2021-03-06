/*
Given an array of distinct integers arr, find all pairs of elements with the minimum absolute difference of any two elements.

Return a list of pairs in ascending order(with respect to pairs), each pair [a, b] follows

a, b are from arr
a < b
b - a equals to the minimum absolute difference of any two elements in arr


Example 1:

Input: arr = [4,2,1,3]
Output: [[1,2],[2,3],[3,4]]
Explanation: The minimum absolute difference is 1. List all pairs with difference equal to 1 in ascending order.
Example 2:

Input: arr = [1,3,6,10,15]
Output: [[1,3]]
Example 3:

Input: arr = [3,8,-10,23,19,-4,-14,27]
Output: [[-14,-10],[19,23],[23,27]]


Constraints:

2 <= arr.length <= 10^5
-10^6 <= arr[i] <= 10^6

*/

package main

import (
	"log"
	"math"
	"sort"
)

func main() {
	tests := [][]int{{4, 2, 1, 3}, {1, 3, 6, 10, 15}, {3, 8, -10, 23, 19, -4, -14, 27}}
	for _, test := range tests {
		log.Printf("minimumAbsDifference(%v) = %v\n", test, minimumAbsDifference(test))

	}
}

// find every unique pair
// make map of [pair diff] -> [pairs list]
// find minimum diff, and return corresponding pairs list
func minimumAbsDifference(arr []int) [][]int {
	diffPairs := map[int][][]int{}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			diff := int(math.Abs(float64(arr[i] - arr[j])))

			_, ok := diffPairs[diff]

			pair := []int{arr[i], arr[j]}
			sort.Ints(pair)

			if !ok {
				// new difference
				diffPairs[diff] = [][]int{pair}
			} else {
				diffPairs[diff] = append(diffPairs[diff], pair)
			}
		}
	}

	// find minimum diff
	var minDiff int

	i := 0
	for diff, _ := range diffPairs {
		if i == 0 {
			minDiff = diff
			i++
		} else if diff < minDiff {
			minDiff = diff
		}
	}

	// return list of pairs associated with minimum diff
	return diffPairs[minDiff]
}
