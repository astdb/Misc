/*
Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.

Example 1:

Input: [3,0,1]
Output: 2
Example 2:

Input: [9,6,4,2,3,5,7,0,1]
Output: 8
Note:
Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?
*/

package main

import (
	"log"
)

func main() {
	tests := [][]int{{0}, {3, 0, 1}, {9, 6, 4, 2, 3, 5, 7, 0, 1}}

	for _, test := range tests {
		log.Printf("missingNumber(%v) == %d\n", test, missingNumber(test))
	}
}

func missingNumber(nums []int) int {
	// v0.1: insert all nums elements into a hashmap (keyed by element).
	// Iterate through 0-len(num), and check if the key is in hashmap.
	elemMap := map[int]bool{}

	for _, i := range nums {
		_, keyExists := elemMap[i]
		if keyExists {
			log.Fatalf("missingNumber(): key already found on element map: %d\n", i)
		} else {
			elemMap[i] = true
		}

	}

	missingNum := 0
	// for j := 0; j < len(nums); j++ {
	for j := 1; j <= len(nums); j++ {
		_, numExists := elemMap[j]
		if !numExists {
			return j
		}
	}

	return missingNum
}
