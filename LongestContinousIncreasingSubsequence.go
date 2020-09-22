/*
Given an unsorted array of integers, find the length of longest continuous increasing subsequence (subarray).

Example 1:
Input: [1,3,5,4,7]
Output: 3
Explanation: The longest continuous increasing subsequence is [1,3,5], its length is 3.
Even though [1,3,5,7] is also an increasing subsequence, it's not a continuous one where 5 and 7 are separated by 4.

Example 2:
Input: [2,2,2,2,2]
Output: 1
Explanation: The longest continuous increasing subsequence is [2], its length is 1.

Note: Length of the array will not exceed 10,000.
*/

package main

import (
	"log"
)

func main() {
	tests := [][]int{{}, {1, 3, 5, 4, 7}, {2, 2, 2, 2, 2}, {2,0,3,4,9,4,9,7,3,4}}

	for _, test := range tests {
		log.Printf("findLengthOfLCIS(%v) = %d\n", test, findLengthOfLCIS(test))
	}
}

func findLengthOfLCIS(nums []int) int {
	lcisLen := 0    // universal longest continuously-increasing subsequence length placeholder
	curLCISLen := 1 // universal longest continuously-increasing subsequence length seen so far

	for i := 0; i < len(nums); i++ {
		if (i - 1) >= 0 { // for indices > 0
			if (nums[i] <= nums[i-1]) || (i == (len(nums) - 1)) {
				if curLCISLen > lcisLen {
					lcisLen = curLCISLen
					curLCISLen = 1
				}
			} else {
				curLCISLen++
			}
		}
	}

	return lcisLen
}
