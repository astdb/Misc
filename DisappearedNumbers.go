/*
Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

Find all the elements of [1, n] inclusive that do not appear in this array.

Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

Example:

Input:
[4,3,2,7,8,2,3,1]

Output:
[5,6]

*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	tests := [][]int{{4, 3, 2, 7, 8, 2, 3, 1}}

	for _, test := range tests {
		fmt.Println(findDisappearedNumbers(test))
	}
}

func findDisappearedNumbers(nums []int) []int {
	sort.Ints(nums)
	absentVals := []int{}

	expectedVal := 1
	prevSeenExpected := 0
	i := 0
	for i < len(nums) {
		if nums[i] == expectedVal {
			prevSeenExpected = expectedVal
			expectedVal++

			if i < len(nums)-1 && nums[i+1] == expectedVal {
				// repeat - skip that
				expectedVal++
				i += 2
			}
		} else {
			for j := prevSeenExpected; j <= expectedVal; j++ {
				absentVals = append(absentVals, j)
			}

			prevSeenExpected = expectedVal
			expectedVal++
			i++
		}
	}

	return absentVals
}
