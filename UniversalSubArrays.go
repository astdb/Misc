
/*
Given an array consisting of 4's and 2's, output how many universal sub-arrays can be formed. 

A universal sub-array is one that consists of a contiguous set of 2's and then a similar-sized contiguous set of 4's (or vice versa). e.g {2,4}, {4,2}, {4,4,4,2,2,2}, {2,2,4,4}

Solution: countingUniversalSubarrays() utilizes helper function getSubArrays() to generate all subarrays of a given array. It then uses another helper function isUniversal() to determine which of those subarrays are universal. A count is made and returned as the result from countingUniversalSubarrays(). 
*/

package main

import (
	"log"
)

func main() {
	tests := [][]int32{{1,2,3,4}, {4,4,2,2,4,2}}

	for _, test := range tests {
		log.Printf("countingUniversalSubarrays(%v) == %v\n", test, countingUniversalSubarrays(test))
	}

func countingUniversalSubarrays(arr []int32) int32 {
	// Write your code here
	
	// compute all subarrays of arr
	res := NewRes()
	getSubArrays(arr, 0, 0, res)

	// check which of arr's subarrays meet universal conditions
	var universalCount int32	// no. of universal subarrays found
	universalCount = 0
	for _, arr := range res.Result {
		if isUniversal(arr) {
			universalCount++
		}
	}

	return universalCount
}

// getSubArrays() recursively computes all sub arrays of a given array, and stores results in result struct property (passed in by reference)
func getSubArrays(arr []int32, start, end int, res *Res) {
	if end >= len(arr) {
		return

	} else if start > end {
		getSubArrays(arr, 0, end + 1, res)

	} else {
		thisSubArr := []int32{}

		for i := start; i < end; i++ {
			thisSubArr = append(thisSubArr, arr[i])
		}

		thisSubArr = append(thisSubArr, arr[end])
		res.Result = append(res.Result, thisSubArr)

		getSubArrays(arr, start + 1, end, res)
	}
}

// Res provides a wrapper for the result storing all subarrays, to be passed by reference to subarray computing function
type Res struct {
	Result [][]int32
}

// NewRes() creates and returns a pointer to an instance of subarray results wrapper struct
func NewRes() *Res {
	var x Res
	return &x
}

// isUniversal() indicates if a given array has universal properties (e.g. 2's and 4's in separate contiguous blocks)
func isUniversal(arr []int32) bool {
	if len(arr) <= 1 {
		return false	// array must have at least two elements to be universal
	}

	if len(arr) % 2 != 0 {
		return false	// array must have an even number of elements to be universal
	}

	twoCount := 0		// no of 2's observed
	fourCount := 0		// no of 4's observed
	switched := false	// flag indicating if switched to 2's from 4's or vice-versa
	inTwos := false		// flag indicating if in 2's section of subarray
	inFours := false	// flag indicating if in 4's section of subarray

	for i := 0; i < len(arr); i++ {
		if !(arr[i] == 2 || arr[i] == 4) {
			return false	// array must contain 2's and 4's only
		}

		if i == 0 {
			// initialize flags and section markers in first element
			if arr[i] == 2 {
				inTwos = true
				twoCount++
			} else if arr[i] == 4 {
				inFours = true
				fourCount++
			} else {
				return false
			}
		} else {
			if inTwos && arr[i] == 2 {
				twoCount++		// 2 encountered while in 2's section of array

			} else if inFours && arr[i] == 4 {
				fourCount++		// 4 encountered while in 4's section of array

			} else if inTwos && arr[i] == 4 && !switched {
				// switching from 2's section to 4's section
				inTwos = false
				inFours = true
				switched = true
				fourCount++

			} else if inFours && arr[i] == 2 && !switched {
				// switching from 4's section to 2's section
				inFours = false
				inTwos = true
				switched = true
				twoCount++

			} else {
				// invalid element or order
				return false
			}
		}
	}

	if twoCount == fourCount {
		return true		// same numbers of 2's and 4's found in contiguous groups
	}

	return false
}
