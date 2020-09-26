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

https://leetcode.com/problems/longest-continuous-increasing-subsequence/submissions/
*/

package main

import (
	"log"
)

func main() {
	// tests := [][]int{{}, {1, 3, 5, 4, 7}, {2, 2, 2, 2, 2}, {2,0,3,4,9,4,9,7,3,4}, {1,3,5,7}, {1,2,3,1,2,3,4,5,6}, {1,2,1}, {1,2,1,1,1,2,3}}
	tests := [][]int{{1,3,5,7}}

	for _, test := range tests {
		log.Printf("findLengthOfLCIS(%v) = %d\n\n", test, findLengthOfLCIS(test))
	}
}

func findLengthOfLCIS(nums []int) int {
	lcisLen := 0    	// overall longest continuously-increasing subsequence (CIS) length
	curLCISLen := 1 	// longest CIS length seen so far

	// for each element in the initial array
	for i := 0; i < len(nums); i++ {

		// for indices > 0	(we will be comparing nums[i] to nums[i-1], so it's important each index we consider has a valid previous index)
		// if (i - 1) >= 0 {
		if i > 0 {

			// if the current element is not greater than the prev element, or if the current index is the last valid index of the nums array.. 
			if (nums[i] <= nums[i-1]) || (i == (len(nums) - 1)) {
				// .. it means we're at the end of a continuously-increasing subsequence.
				if i == (len(nums) - 1) {
					log.Printf("\t\tfindLengthOfLCIS(): end-of-array\n")
				}

				// if the length of the current longest CIS is less than the one we saw just now, it needs to be updated. 
				if curLCISLen > lcisLen {
					if i == (len(nums) - 1) {
						log.Printf("\ti == (len(nums) - 1)")
					}
					log.Printf("\tUpdating lcisLen - index %d\n", i)
					lcisLen = curLCISLen
					curLCISLen = 1
				}
			} else {
				curLCISLen++
			}
		}
	}

	// if curLCISLen > lcisLen {
	// 	lcisLen = curLCISLen
		
	// }
	return lcisLen
}
