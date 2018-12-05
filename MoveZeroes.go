/*
Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Example:

Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
Note:

You must do this in-place without making a copy of the array.
Minimize the total number of operations.
*/

package main

import (
	"fmt"
)

func main() {
	tests := [][]int{{}, {0}, {0,1}, {1,0}, {0,0,1}, {0,0,0,1}, {0, 1, 0, 3, 12}}

	for _, test := range tests {
		fmt.Println("\nInput:", test)
		moveZeroes(test)
		fmt.Println("Output:", test)
	}
}

// func moveZeroes(nums []int) {
// 	end := len(nums)
// 	for i := 0; i < end; i++ {

// 	}
// }

func moveZeroes(nums []int) {
	end := len(nums) - 1
	for i := 0; i < end; {
		if nums[i] == 0 {
			j := i
			for j+1 <= end {
				nums[j] = nums[j+1]
				j++
			}

			nums[end] = 0
			end--
		} else {
			i++

		}
	}
}
