/*
Given an array of integers, find if the array contains any duplicates.

Your function should return true if any value appears at least twice in the array, and it should return false if every element is distinct.

Example 1:

Input: [1,2,3,1]
Output: true
Example 2:

Input: [1,2,3,4]
Output: false
Example 3:

Input: [1,1,1,3,3,4,3,2,4,2]
Output: true
*/

package main

import (
	"fmt"
	"sort"
	"math/rand"
	"time"
)

func main()  {
	// tests := [][]int{{}, {1},{1,2,3,1}, {1,2,3,4}, {1,1,1,3,3,4,3,2,4,2}}

	// for _, test := range tests {
	// 	fmt.Println(test, containsDuplicate2(test))
	// }

	// test function performances against large input
	bigArray := []int{}
	for i := 0; i < 1000000; i++ {
		bigArray = append(bigArray, rand.Intn(1000))
	}

	fmt.Println("Running sort-based function..")
	start1 := time.Now()
	fmt.Println(containsDuplicate1(bigArray))
	end1 := time.Now()
	fmt.Println("Elapsed time(ms):", end1.Sub(start1))

	fmt.Println("\nRunning map-based function..")
	start2 := time.Now()
	fmt.Println(containsDuplicate2(bigArray))
	end2 := time.Now()
	fmt.Println("Elapsed time(ms):", end2.Sub(start2))
}

// sort-based approach (space efficient)
func containsDuplicate1(nums []int) bool {
	// sort input
	sort.Ints(nums)	

	// check for consequent duplicate elements
	for i := 0; i <= len(nums)-2; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}

// map-based approach (runtime efficient)
func containsDuplicate2(nums []int) bool {
	// declare map to store element counts
	m := map[int]int{}

	// per element in input
	for _, i := range nums {
		// check if in map already, if yes: duplicate
		_, exists := m[i]
		if exists {
			return true
		} else {
			// if not in map, insert
			m[i] = 1
		}
	}

	return false
}
