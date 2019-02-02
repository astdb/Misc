/*
Given an array A of integers, for each integer A[i] we may choose any x with -K <= x <= K, and add x to A[i].
After this process, we have some array B.
Return the smallest possible difference between the maximum value of B and the minimum value of B.

Example 1:
Input: A = [1], K = 0
Output: 0
Explanation: B = [1]


Example 2:
Input: A = [0,10], K = 2
Output: 6
Explanation: B = [2,8]


Example 3:
Input: A = [1,3,6], K = 3
Output: 0
Explanation: B = [3,3,3] or B = [4,4,4]

*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	tests := [][][]int{{{1}, {0}}, {{0, 10}, {2}}, {{1, 3, 6}, {3}}, {{2, 7, 2}, {1}}}

	for _, test := range tests {
		fmt.Printf("%d\n", smallestRangeI(test[0], test[1][0]))
	}
}

func smallestRangeI(A []int, K int) int {
	// sort slice (to find median)
	sort.Ints(A)
	// find median
	var median int
	if len(A)%2 == 0 {
		// even length array
		median = (A[len(A)/2] + A[len(A)/2-1]) / 2
	} else {
		// odd length array
		if len(A) == 1 {
			median = A[0]
		} else {
			median = A[len(A)/2]
		}
	}

	fmt.Printf("-------------------\nA: %v\nK: %d\n", A, K)
	fmt.Printf("\nMedian:%d\n", median)

	var smallest int
	var largest int

	for k, v := range A {
		fmt.Printf("\n\tIndex: %d, Value: %d\n", k, v)
		var newV int
		if v > median {
			// need to add a negative amount
			fmt.Println("\tValue greater than median, need to be reduced..")
			if v-median < K {
				// v can be brought down to median value
				newV = median
				fmt.Printf("\tValue can be brought down to median (%d)\n", newV)
			} else {
				// this is the closest we can bring v to median
				newV = v - K
				fmt.Printf("\tClosest value can be brought down to median is %d\n", newV)
			}
		} else if v < median {
			// need to add a positive amount
			fmt.Println("\tValue less than median, need to be increased..")
			if median-v < K {
				// v can be brought up to median value
				newV = median
				fmt.Printf("\tValue can be brought up to median (%d)\n", newV)
			} else {
				// closest v can be brought to median
				newV = v + K
				fmt.Printf("\tClosest value can be brought up to median is %d\n", newV)
			}
		} else {
			// v == median
			fmt.Printf("\tMedian (%d) == Value (%d), no increment/decrement required.", median, v)
			newV = v
		}

		// keep track of largest and smallest newV values
		if k == 0 {
			smallest = newV
			largest = newV
		} else {
			if newV < smallest {
				smallest = newV
			}

			if newV > largest {
				largest = newV
			}
		}

		fmt.Printf("\n\tCurrent Largest: %d\n\tCurrent Smallest: %d\n", largest, smallest)
	}

	return largest - smallest
}
