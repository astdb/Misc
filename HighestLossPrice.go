/*
Problem: Given a list of daily prices, calculate the largest loss that can be incurred on buying an item on any day and selling  on any day afterwards.

Solution: An initial brute-force version of the solution would be for to traverse the list of prices, and for each day traverse the rest of the list to see the lowest future price. This would provide the correct result but would run at exponential runtime complexity.

A better version of the solution depends on the insight that for each day, the program only needs to know the highest past price observed. This can then be used to calculate the potential loss for each day, and the highest loss would be the sought after result. This can be calculated in a single pass of the price list with linear runtime complexity and constant space complexity. Below implementation uses this method.
*/

package main

import (
	"log"
)

func main() {
	tests := [][]int{{7, 4, 2, 9}, {2, 4, 6, 8}, {7, 4, 8, 2, 9}}

	for i, test := range tests {
		log.Printf("Test #%d: worstLosingStreak(%v) == %d\n", (i + 1), test, worstLosingStreak(test))
	}
}

func worstLosingStreak(nums []int) int {
	highestPrice := 0 // highest price observed so far
	highestLoss := 0  // highest loss observed so far

	// for each price
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			// initialize highest price observed so far to first day's price
			highestPrice = nums[i]
		} else {
			// if selling at today's price relative to the highest past price yields the largest loss, update highest loss so far
			if (highestPrice - nums[i]) > highestLoss {
				highestLoss = highestPrice - nums[i]
			}

			// if today's price is the highest seen so far, update highest price so far
			if nums[i] > highestPrice {
				highestPrice = nums[i]
			}
		}
	}

	return highestLoss
}
