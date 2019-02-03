/*
Given an array of integers and a pivot element index, reorder the array so that elements smaller than the pivot appear first, followed by elements equal to pivot and then elements larger that the pivot (i.e. 'Dutch National Flag' grouping).
*/

package main

import (
	"fmt"
)

func main() {
	tests := [][][]int{{{0, 1, 2, 0, 2, 1, 1}, {3}}, {{0, 1, 2, 0, 2, 1, 1}, {1}}}

	for _, test := range tests {
		// fmt.Printf("DNFGrouping(%v, %d) == ", test[0], test[1][0])
		// fmt.Printf("%v\n", DNFGrouping(test[0], test[1][0]))
		DNFGrouping(test[0], test[1][0])
	}
}

func DNFGrouping(nums []int, pivotIndex int) []int {
	if len(nums) <= 1 {
		// nothing to sort
		return nums
	}

	fmt.Printf("nums: %v, PivotIndex: %d\n", nums, pivotIndex)

	// two- or more element array
	if pivotIndex < len(nums) {
		pivot := nums[pivotIndex]

		// indexes where lower than pivot, equal to pivot, higher than the pivot, and unsorted sections of the array start
		// lowStart := 0
		eqStart := 0
		highStart := 0
		unSortedStart := 0

		for i := 0; i < len(nums); i++ {
			if i == 0 {
				// start by moving pivot to the start of array
				swap(nums, 0, pivotIndex)
				pivotIndex = 0
			}

			if nums[i] < pivot {
				// everything from the start of the equal section must move one position to the right, and nums i swapped
				fmt.Printf("\n------------------\nnums[%d] (%d) is < pivot (%d)\n", i, nums[i], pivot)
				fmt.Printf("Before move: eqStart: %d, highStart: %d, unSortedStart: %d\nMoving elements...\n", eqStart, highStart, unSortedStart)
				tmp := nums[i]
				j := i
				for j-1 >= eqStart {
					nums[j] = nums[j-1]
					j--
				}

				nums[eqStart] = tmp
				eqStart++
				pivotIndex++
				unSortedStart++
				i = unSortedStart

				fmt.Printf("After move: eqStart: %d, highStart: %d, unSortedStart: %d\nnums: %v\n", eqStart, highStart, unSortedStart, nums)

				if unSortedStart >= len(nums) {
					// everything's sorted
					break
				}
			}

			if nums[i] == pivot {
				// everything from the start of the higher section must move one position to the right, and nums i swapped
				fmt.Printf("\n------------------\nnums[%d] (%d) is == pivot (%d)\n", i, nums[i], pivot)
				fmt.Printf("Before move: eqStart: %d, highStart: %d, unSortedStart: %d\nMoving elements...\n", eqStart, highStart, unSortedStart)

				tmp := nums[i]
				j := i
				for j-1 >= highStart {
					nums[j] = nums[j-1]
					j--
				}

				nums[highStart] = tmp
				highStart++
				// pivotIndex++
				unSortedStart++
				i = unSortedStart

				fmt.Printf("After move: eqStart: %d, highStart: %d, unSortedStart: %d\nnums: %v\n", eqStart, highStart, unSortedStart, nums)

				if unSortedStart >= len(nums) {
					// everything's sorted
					break
				}

			}

			if nums[i] > pivot {
				// everything from the start of the unsorted section must move one position to the right, and nums i swapped
				fmt.Printf("\n------------------\nnums[%d] (%d) is > pivot (%d)\n", i, nums[i], pivot)
				fmt.Printf("Before move: eqStart: %d, highStart: %d, unSortedStart: %d\nMoving elements...\n", eqStart, highStart, unSortedStart)

				tmp := nums[i]
				j := i
				for j-1 >= unSortedStart {
					nums[j] = nums[j-1]
					j--
				}

				nums[unSortedStart] = tmp
				unSortedStart++
				i = unSortedStart
				// pivotIndex++
				// unsortedStart++

				fmt.Printf("After move: eqStart: %d, highStart: %d, unSortedStart: %d\nnums: %v\n", eqStart, highStart, unSortedStart, nums)

				if unSortedStart >= len(nums) {
					// everything's sorted
					break
				}

			}
		}
	}

	return nums
}

func swap(nums []int, x int, y int) {
	nums[x], nums[y] = nums[y], nums[x]
}
