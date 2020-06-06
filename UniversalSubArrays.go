
/*
Given an array consisting of 4's and 2's, output how many universal sub-arrays can be formed. 

A universal sub-array is one that consists of a contiguous set of 2's and then a similar-sized contiguous set of 4's (or vice versa). e.g {2,4}, {4,2}, {4,4,4,2,2,2}, {2,2,4,4}
*/

package main

import (
	"log"
	// "fmt"
)

func main() {
	tests := [][]int32{{1,2,3,4}, {4,4,2,2,4,2}}

	for _, test := range tests {
		log.Printf("countingUniversalSubarrays(%v) == %v\n", test, countingUniversalSubarrays(test))
		// log.Printf();
	}

	// res := [][]int32{}
	// printSubArrays([]int32{1,2,3,4}, 0, 0, res)

	// fmt.Println(res)

	// x := []int{}

	// res := NewRes()
	// getSubArrays([]int32{1,2,3,4}, 0, 0, res)
	// fmt.Println(res.Result)

	// tests := [][]int32{{4,2}, {2,4}, {4,4,2,2}, {2,2,4,4}, {4,4,4,2,2,2}, {2}, {4}, {2,2,4}, {4,2,4}, {4,5,2,2}}

	// for _, test := range tests {
	// 	fmt.Printf("isUniversal(%v) == %v\n", test, isUniversal(test))
	// }
}

func testFunc(x []int) {
	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)
	x = append(x, 4)
}

func countingUniversalSubarrays(arr []int32) int32 {
    // Write your code here
	res := NewRes()
	// getSubArrays([]int32{1,2,3,4}, 0, 0, res)
	getSubArrays(arr, 0, 0, res)

	var universalCount int32
	universalCount = 0
	for _, arr := range res.Result {
		if isUniversal(arr) {
			universalCount++
		}
	}

	return universalCount
}

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

type Res struct {
	Result [][]int32
}

func NewRes() *Res {
	var x Res
	return &x
}

func isUniversal(arr []int32) bool {
	if len(arr) <= 1 {
		return false
	}

	if len(arr) % 2 != 0 {
		return false
	}

	twoCount := 0
	fourCount := 0
	switched := false
	inTwos := false
	inFours := false

	for i := 0; i < len(arr); i++ {
		if !(arr[i] == 2 || arr[i] == 4) {
			return false
		}

		if i == 0 {
			if arr[i] == 2 {
				inTwos = true
				twoCount++
			} else if arr[i] == 4 {
				inFours = true
				fourCount++
			} else {
				log.Printf("\tisUniversal(): i == 0, unexpected arr val, returning false")
				return false
			}
		} else {
			if inTwos && arr[i] == 2 {
				twoCount++
			} else if inFours && arr[i] == 4 {
				fourCount++
			} else if inTwos && arr[i] == 4 && !switched {
				inTwos = false
				inFours = true
				switched = true
				fourCount++
			} else if inFours && arr[i] == 2 && !switched {
				inFours = false
				inTwos = true
				switched = true
				twoCount++
			} else {
				return false
			}
		}
	}

	if twoCount == fourCount {
		return true
	}

	return false
}
