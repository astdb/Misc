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

func containsDuplicate1(nums []int) bool {
	sort.Ints(nums)	
	for i := 0; i <= len(nums)-2; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}

func containsDuplicate2(nums []int) bool {
	m := map[int]int{}

	for _, i := range nums {
		_, exists := m[i]
		if exists {
			return true
		} else {
			m[i] = 1
		}
	}

	return false
}
