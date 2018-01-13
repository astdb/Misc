
/*
    Given a non-empty integer array of size n, find the minimum number of moves required to make all array elements equal, where a move is incrementing (n-1) elements by 1.
*/

package main

import (
	"fmt"
)

func main() {
	tests := [][]int{{}, {1}, {1,1}, {1,2,3,}, {4,5,6,}, {7,8,9,}, {1,2147483648}}

	for _, testcase := range tests {
		fmt.Printf("%v\t%d\n", testcase, minMoves(testcase))
	}
}

func minMoves(nums []int) int {
	// while not all array elements are equal
	// fmt.Printf("\n-------------------\nStarting array: %v\n", nums)
	moves := 0
    for !equalElements(nums) {
		moves = moves + 1
		highestIndex := getHighestIndex(nums)

		for i := 0; i < len(nums); i++ {
			if i != highestIndex {
				nums[i] = nums[i] + 1
			}
		}

		// fmt.Printf("\tMove %d, Array: %v\n", moves, nums)
	}

	return moves
}

func getHighestIndex(n []int) int {
	if len(n) <= 0 {
		return 0
	}

	highestIndex := 0
	for k,v := range n {
		if v > n[highestIndex] {
			highestIndex = k
		}
	}

	return highestIndex
}

func equalElements(n []int) bool {
	if len(n) <= 0 {
		return true
	}

	k := n[0]
	for _, v := range n {
		if k != v {
			return false
		}
	}

	return true
}
