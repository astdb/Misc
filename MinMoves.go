
/*
    Given a non-empty integer array of size n, find the minimum number of moves required to make all array elements equal, where a move is incrementing n - 1 elements by 1.
*/

package main

import (
	"fmt"
)

func main() {
	tests := [][]int{{1,2,3,}, {4,5,6,}, {7,8,9,}}

	for _, testcase := range tests {
		fmt.Printf("%v\t%d", testcase, minMoves(testcase))
	}
}

func minMoves(nums []int) int {
	// while not all array elements are equal
	moves := 0
    for !equalElements(nums) {
		moves = moves + 1
		highest := getHighest(nums)

		for i := 0; i < len(nums); i++ {
			if i != highest {
				nums[i] = nums[i] + 1
			}
		}
	}

	return moves
}

func getHighest(n []int) int {
	highest := n[0]

	for _, v := range n {
		if v > highest {
			highest = v
		}
	}

	return highest
}

func equalElements(n []int) bool {
	if len(n) <= 0 {
		return false
	}

	k := n[0]
	for _, v := range n {
		if k != v {
			return false
		}
	}

	return true
}
