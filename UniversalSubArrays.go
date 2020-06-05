
/*
Given an array consisting of 4's and 2's, output how many universal sub-arrays can be formed. 

A universal sub-array is one that consists of a contiguous set of 2's and then a similar-sized contiguous set of 4's (or vice versa). e.g {2,4}, {4,2}, {4,4,4,2,2,2}, {2,2,4,4}
*/

package main

import (
	// "log"
	"fmt"
)

func main() {
	// tests := [][]int{{1,2,3,4}, {4,4,2,2,4,2}}

	// for _, test := range tests {
	// 	// log.Printf("countingUniversalSubarrays(%v) == %v\n", test, countingUniversalSubarrays(test))
	// 	log.Printf();
	// }

	// res := [][]int32{}
	// printSubArrays([]int32{1,2,3,4}, 0, 0, res)

	// fmt.Println(res)

	// x := []int{}

	res := NewRes()
	printSubArrays([]int32{1,2,3,4}, 0, 0, res)
	fmt.Println(res.Result)
}

func testFunc(x []int) {
	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)
	x = append(x, 4)
}

// func countingUniversalSubarrays(arr []int32) int32 {
//     // Write your code here
// 	start := 0
// 	end := 0

// 	if end >= len(arr) {
// 		return 0
// 	} else if start > end {

// 	}
// }

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

	twoSect := false
	sectorChanged := false

	twoCount := 0
	fourCount := 0

	for i := 0; i < len(arr); i++ {
		if i == 0 {
			if arr[i] == 2 {
				twoSect = true
				twoCount++
			} else if arr[i] == 4 {
				twoSect = true
				fourCount++
			} else {
				return false
			}
		}

		if arr[i] == 2 && twoSect {
			twoCount++
		}

		if arr[i] == 4 && fourSect {
			fourCount++
		}


	}
}

func getSubArrays(arr []int32, start, end int, res *Res) {
	// start := 0
	// end := 0

	if end >= len(arr) {
		return
	} else if start > end {
		printSubArrays(arr, 0, end+1, res)
	} else {
		thisArr := []int32{}
		// fmt.Printf("[")

		for i := start; i < end; i++ {
			// fmt.Printf("%d, ", arr[i])
			thisArr = append(thisArr, arr[i])
		}

		// fmt.Printf("%d]", arr[end])
		thisArr = append(thisArr, arr[end])
		res.Result = append(res.Result, thisArr)

		printSubArrays(arr, start+1, end, res)
	}
}
