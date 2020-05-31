/*
One approach would be to iterate through the set, pushing each value onto a map (as a key). This would reveal the duplicate item. Next, each value within the valid range for the set can be checked in the keys of the map. This should reveal the missing item.

Alternate approach would be to sort the set of values and find the duplicate and missing items.
*/

package main

import (
	"log"
)

func main() {
	tests := [][]int{{1, 2, 2, 4}}

	for _, test := range tests {
		log.Printf("findMismatch(%v) == %v\n", test, findMismatch(test))
	}
}

func findMismatch(nums []int) []int {
	// vars holding duplicate and missing valuee in set
	dup := 0
	mis := 0

	// insert all values in set to map as keys
	numsMap := map[int]int{}

	for _, val := range nums {
		// log.Printf("Inserting %d as key to numsMap..\n", val)
		_, exists := numsMap[val]
		if exists {
			// log.Printf("%d exists on numsMap - recording as duplicate..\n", val)
			// if a set value already exists in map as a key
			// that would be the repeated value.
			dup = val
		} else {
			numsMap[val] = 1
			// log.Printf("%d not in map - added now numMap () \n")
		}
	}

	// iterate through all valid values for set
	// if any doesn't exist
	for i := 1; i <= len(nums); i++ {
		_, exists := numsMap[i]
		if !exists {
			mis = i
		}
	}

	return []int{dup, mis}
}
