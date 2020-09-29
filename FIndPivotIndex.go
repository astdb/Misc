/*
Given an array of integers nums, write a method that returns the "pivot" index of this array.

We define the pivot index as the index where the sum of all the numbers to the left of the index is equal to the sum of all the numbers to the right of the index.

If no such index exists, we should return -1. If there are multiple pivot indexes, you should return the left-most pivot index.

 

Example 1:

Input: nums = [1,7,3,6,5,6]
Output: 3
Explanation:
The sum of the numbers to the left of index 3 (nums[3] = 6) is equal to the sum of numbers to the right of index 3.
Also, 3 is the first index where this occurs.
Example 2:

Input: nums = [1,2,3]
Output: -1
Explanation:
There is no index that satisfies the conditions in the problem statement.
 

Constraints:

The length of nums will be in the range [0, 10000].
Each element nums[i] will be an integer in the range [-1000, 1000].
*/

package main

import (
	"log"
)

func main() {
	tests := [][]int{{1,7,3,6,5,6}, {1,2,3}}
	for _, test := range tests {
		log.Printf("pivotIndex(%v) == %d\n", test, pivotIndex(test))
	}
}

func pivotIndex(nums []int) int {
		// totals from left
		totalsLeft := []int{}

		total := 0
		for i := 0; i < len(nums); i++ {
			total += nums[i]
			totalsLeft = append(totalsLeft, total)
		}

		log.Printf("pivotIndex(): totalsLeft: %v\n", totalsLeft)

		// totals from right
		totalsRight := []int{}

		total = 0
		for i := len(nums)-1; i >= 0; i-- {
			total += nums[i]
			totalsRight = append(totalsRight, total)
		}

		log.Printf("pivotIndex(): totalsRight: %v\n", totalsRight)

		for i := 0; i < len(totalsLeft); i++ {
			leftTot := 0
			if i-1 >= 0 {
				leftTot = totalsLeft[i-1]

				rightTot := 0
				
				if (len(totalsRight)-(i+2)) >= 0 {
					rightTot = totalsRight[len(totalsRight)-(i+2)]

					if leftTot == rightTot {
						return i
					}		
				}	
			}

			// rightTot := 0
			// if i+1 < len(totalsRight) {
			// 	rightTot = totalsRight[i+1]
			// }

			// if (len(totalsRight)-(i+2)) >= 0 {
			// 	rightTot = totalsRight[len(totalsRight)-(i+2)]
			// }

			// if leftTot == rightTot {
			// 	return i
			// }
		}

		return -1
}
