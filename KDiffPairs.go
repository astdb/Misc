/*
Given an array of integers and an integer k, you need to find the number of unique k-diff pairs in the array. Here a k-diff pair is defined as an integer pair (i, j), where i and j are both numbers in the array and their absolute difference is k.

Example 1:
Input: [3, 1, 4, 1, 5], k = 2
Output: 2
Explanation: There are two 2-diff pairs in the array, (1, 3) and (3, 5).
Although we have two 1s in the input, we should only return the number of unique pairs.

Example 2:
Input:[1, 2, 3, 4, 5], k = 1
Output: 4
Explanation: There are four 1-diff pairs in the array, (1, 2), (2, 3), (3, 4) and (4, 5).

Example 3:
Input: [1, 3, 1, 5, 4], k = 0
Output: 1
Explanation: There is one 0-diff pair in the array, (1, 1).

Note:
The pairs (i, j) and (j, i) count as the same pair.
The length of the array won't exceed 10,000.
All the integers in the given input belong to the range: [-1e7, 1e7].
*/

// NOTE: THIS IS O(N^2) - OPTIMIZE

package main

import (
	"log"
)

func main() {
	tests := [][][]int{{{3, 1, 4, 1, 5}, {2}}, {{1, 2, 3, 4, 5}, {1}}, {{1, 3, 1, 5, 4}, {0}}}

	for _, test := range tests {
		log.Printf("findPairs(%v, %d) == %d\n", test[0], test[1][0], findPairs(test[0], test[1][0]))
	}
}

func findPairs(nums []int, k int) int {
	pairCount := 0
	pairsFound := [][]int{}

	for i := 0; i < len(nums); i++ {
		curElm := nums[i]

		for j := i + 1; j < len(nums); j++ {
			if abs(curElm-nums[j]) == k {
				thisPair := []int{curElm, nums[j]}
				// log.Println(thisPair)
				if !pairSeen(thisPair, pairsFound) {
					pairsFound = append(pairsFound, thisPair)
					pairCount++
				}

			}
		}
	}

	return pairCount
}

func pairSeen(pair []int, pairsFound [][]int) bool {
	// log.Printf("\tlen(pairsFound) == %d\n", len(pairsFound))
	for _, thisPairFound := range pairsFound {
		if (pair[0] == thisPairFound[0] && pair[1] == thisPairFound[1]) || (pair[0] == thisPairFound[1] && pair[1] == thisPairFound[0]) {
			return true
		}
	}

	return false
}

func abs(x int) int {
	if x < 0 {
		return x * (-1)
	}

	return x
}
