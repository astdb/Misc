/*
Given an array arr of positive integers sorted in a strictly increasing order, and an integer k.

Find the kth positive integer that is missing from this array.

Example 1:

Input: arr = [2,3,4,7,11], k = 5
Output: 9
Explanation: The missing positive integers are [1,5,6,8,9,10,12,13,...]. The 5th missing positive integer is 9.

Example 2:

Input: arr = [1,2,3,4], k = 2
Output: 6
Explanation: The missing positive integers are [5,6,7,...]. The 2nd missing positive integer is 6.

Constraints:

1 <= arr.length <= 1000
1 <= arr[i] <= 1000
1 <= k <= 1000
arr[i] < arr[j] for 1 <= i < j <= arr.length

*/

package main

import (
	"log"
)

func main() {
	tests := [][][]int{{{1, 3}, {1}}, {{2, 3, 4, 7, 11}, {5}}, {{1, 2, 3, 4}, {2}}}
	for _, test := range tests {
		log.Printf("findKthPositive(%v, %d) = %d\n", test[0], test[1][0], findKthPositive(test[0], test[1][0]))
	}
}

// given a strictly-increasing positive integer array and another positive int k, return the k-th positive integer missing from the array
func findKthPositive(arr []int, k int) int {
	intCount := 1     // count representing positive integers, starting from 1
	missingCount := 0 // count representing missing positive integers from arr
	arrIndex := 0     // current index on arr we're looking at

	// while we're yet to see the k-th missing int
	for missingCount < k {

		// ensure array index is within arr's length
		if arrIndex < len(arr) {
			if arr[arrIndex] == intCount {
				// not a missing int - increment both int and array index
				intCount++
				arrIndex++

			} else {
				// arr[index] is bigger than current int - increase int until arr[index] == int (and count missing ints)
				if arr[arrIndex] > intCount {
					for intCount < arr[arrIndex] {
						missingCount++
						if missingCount == k {
							return intCount
						}

						intCount++

					}

					// at this point, arr[index] == int
					intCount++
					arrIndex++

				} else {
					// arr[index] < int : cannot happen as arr contents are strictly increasing

				}

			}
		} else {
			for missingCount < k {
				missingCount++
				if missingCount == k {
					return intCount
				}

				intCount++
			}
		}
	}

	return intCount
}
