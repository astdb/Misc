/*
Given an array of positive integers arr, calculate the sum of all possible odd-length subarrays.

A subarray is a contiguous subsequence of the array.

Return the sum of all odd-length subarrays of arr.



Example 1:

Input: arr = [1,4,2,5,3]
Output: 58
Explanation: The odd-length subarrays of arr and their sums are:
[1] = 1
[4] = 4
[2] = 2
[5] = 5
[3] = 3
[1,4,2] = 7
[4,2,5] = 11
[2,5,3] = 10
[1,4,2,5,3] = 15
If we add all these together we get 1 + 4 + 2 + 5 + 3 + 7 + 11 + 10 + 15 = 58
Example 2:

Input: arr = [1,2]
Output: 3
Explanation: There are only 2 subarrays of odd length, [1] and [2]. Their sum is 3.
Example 3:

Input: arr = [10,11,12]
Output: 66


Constraints:

1 <= arr.length <= 100
1 <= arr[i] <= 1000

*/

package main

import (
	"log"
)

func main() {
	tests := [][]int{{1, 4, 2, 5, 3}, {1, 2}}
	for _, test := range tests {
		log.Printf("sumOddengthSubarrays(%v) = %d\n", test, sumOddLengthSubarrays(test))
	}
}

func sumOddLengthSubarrays(arr []int) int {
	r := NewRes()
	getSubArrays(arr, 0, 0, r)

	oddSubArrTot := 0
	for _, subArr := range r.Result {
		if len(subArr)%2 != 0 {
			oddSubArrTot += sumArr(subArr)
		}
	}

	return oddSubArrTot
}

func sumArr(arr []int) int {
	sum := 0

	for _, i := range arr {
		sum += i
	}

	return sum
}

// getSubArrays() recursively computes all sub arrays of a given array, and stores results in result struct property (passed in by reference)
func getSubArrays(arr []int, start, end int, res *Res) {
	if end >= len(arr) {
		return

	} else if start > end {
		getSubArrays(arr, 0, end+1, res)

	} else {
		thisSubArr := []int{}

		for i := start; i < end; i++ {
			thisSubArr = append(thisSubArr, arr[i])
		}

		thisSubArr = append(thisSubArr, arr[end])
		res.Result = append(res.Result, thisSubArr)

		getSubArrays(arr, start+1, end, res)
	}
}

// Res provides a wrapper for the result storing all subarrays, to be passed by reference to subarray computing function
type Res struct {
	Result [][]int
}

// NewRes() creates and returns a pointer to an instance of subarray results wrapper struct
func NewRes() *Res {
	var x Res
	return &x
}
