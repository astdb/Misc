/*
Given an array and a value, remove all instances of that value in-place and return the new length.
Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
The order of elements can be changed. It doesn't matter what you leave beyond the new length.

*/

package main

import (
	"fmt"
)

func main() {

}

func removeElement(nums []int, val int) int {
	top := len(nums)	// size of the new array

	for i := 0; i < top; i++ {
		if i == top-1 && nums[i] == val {
			top--
		} else {
			if nums[i] == val {
			// copy the rest of the array down and overwrite nums[i]
			for j := i; j < top-1; j++ {
				nums[j] = nums[j+1]
			}

			top--
			i--
		}

		}
	}

	return top
}
