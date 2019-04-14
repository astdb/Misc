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
	"fmt"
)

func main() {
	tests := [][][]int{{{-1, 0, 3, 5, 9, 12}, {9}}, {{-1, 0, 3, 5, 9, 12}, {2}}}

	for _, test := range tests {
		fmt.Printf("Searching for %d in %v: result index = %d\n", test[1][0], test[0], search(test[0], test[1][0]))
	}
}

func search(nums []int, target int) int {
	if nums[len(nums)/2] == target {
		return len(nums) / 2
	}

	// if len(nums)/2 <= 0 {
	// 	return -1
	// }

	if nums[len(nums)/2] < target {
		return search(nums[:len(nums)/2], target)
	} else {
		if (len(nums)/2)+1 < len(nums) {
			return search(nums[(len(nums)/2)+1:], target)
		} else {
			return -1
		}		
	}

}
