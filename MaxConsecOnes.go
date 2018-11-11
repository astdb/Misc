/*
Given a binary array, find the maximum number of consecutive 1s in this array.

Example 1:
Input: [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s.
    The maximum number of consecutive 1s is 3.
Note:

The input array will only contain 0 and 1.
The length of input array is a positive integer and will not exceed 10,000
Seen this question in a real interview before?  

*/


package main

import (
    "fmt"
)

func main() {
	testInput := []int{1,1,0,1,1,1}
	fmt.Println(findMaxConsecutiveOnes(testInput))
}

func findMaxConsecutiveOnes(nums []int) int {
	maxLen := 0
	curLen := 0

	for _,v := range nums {
		if v == 1 {
			curLen++
		} else {
			if curLen > maxLen {
				maxLen = curLen
			}

			curLen = 0
		}
	}

	if curLen > maxLen {
		return curLen
	}

	return maxLen
}
