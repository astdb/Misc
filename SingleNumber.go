/*
Given a non-empty array of integers, every element appears twice except for one. Find that single one.

Note:

Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Example 1:

Input: [2,2,1]
Output: 1
Example 2:

Input: [4,1,2,1,2]
Output: 4
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	tests := [][]int{{2, 2, 1}, {4, 1, 2, 1, 2}}
	// fmt.Println(singleNumber([]int{2, 2, 1}))
	for _, test := range tests {
		fmt.Println(SingleNumber2(test))
	}
}

func SingleNumber2(nums []int) int {
	sort.Ints(nums)

	i := 0
	for i < len(nums) {
		if i == len(nums)-1 {
			return nums[i]
		}

		if i+1 < len(nums) {
			if nums[i] != nums[i+1] {
				return nums[i]
			} else {
				i += 2
			}
		}
	}

	return i
}

func singleNumber(nums []int) int {
	ones := 0
	twos := 0
	var common_bit_mask int

	for _, v := range nums {
		twos = twos | (ones & v)
		ones = ones ^ v
		common_bit_mask = ^(ones & twos)
		ones = ones & common_bit_mask
		twos = twos & common_bit_mask
	}

	return ones
}
