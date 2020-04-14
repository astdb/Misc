/*
Given a sorted (in ascending order) integer array nums of n elements and a target value, write a function to search target in nums. If target exists, then return its index, otherwise return -1.


Example 1:

Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4

Example 2:

Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1
 

Note:

You may assume that all elements in nums are unique.
n will be in the range [1, 10000].
The value of each element in nums will be in the range [-9999, 9999].
*/

package main

import (
	"log"
)

func main() {
	var input []int
	var target int
	var expOutput int
	var output int

	input = []int{-1,0,3,5,9,12}
	target = 9
	expOutput = 4
	output = search(input, target)
	log.Printf("search(%v, %d) == %d", input, target, output)
	if output != expOutput {
		log.Printf("\tFAIL\n")
	} else {
		log.Printf("\tPASS\n")
	}

	input = []int{-1,0,3,5,9,12}
	target = 2
	expOutput = -1
	output = search(input, target)
	log.Printf("search(%v, %d) == %d", input, target, output)
	if output != expOutput {
		log.Printf("\tFAIL\n")
	} else {
		log.Printf("\tPASS\n")
	}
}

// binary search
func search(nums []int, target int) int {
	start := 0
	end := len(nums)
	mid := start + ((end-start)/2)

	for {
		if nums[mid] < target {
			// look in the upper half
			start = mid+1
			mid = start + ((end-start)/2)

		} else if nums[mid] > target {
			// look in the lower half
			end = mid
			mid = start + ((end-start)/2)

		} else if nums[mid] == target {
			return mid
		}

		if start >= end {
			return -1
		}
	}

}

// linear search
func search1(nums []int, target int) int {
	for k, v := range nums {
		if v == target {
			return k
		}
	}

	return -1
}
