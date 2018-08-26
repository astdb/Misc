/*
Alice and Bob have candy bars of different sizes: A[i] is the size of the i-th bar of candy that Alice has, and B[j] is the size of the j-th bar of candy that Bob has.

Since they are friends, they would like to exchange one candy bar each so that after the exchange, they both have the same total amount of candy.  (The total amount of candy a person has is the sum of the sizes of candy bars they have.)

Return an integer array ans where ans[0] is the size of the candy bar that Alice must exchange, and ans[1] is the size of the candy bar that Bob must exchange.

If there are multiple answers, you may return any one of them.  It is guaranteed an answer exists.
 

Example 1:

Input: A = [1,1], B = [2,2]
Output: [1,2]
Example 2:

Input: A = [1,2], B = [2,3]
Output: [1,2]
Example 3:

Input: A = [2], B = [1,3]
Output: [2,3]
Example 4:

Input: A = [1,2,5], B = [2,4]
Output: [5,4]
 

Note:

1 <= A.length <= 10000
1 <= B.length <= 10000
1 <= A[i] <= 100000
1 <= B[i] <= 100000
It is guaranteed that Alice and Bob have different total amounts of candy.

*/

package main

import (
	"fmt"
)

func main() {
	testCases := [][][]int{{{1,1},{2,2}}, {{1,2},{2,3}}, {{2},{1,3}}, {{1,2,5},{2,4}}}

	for _, testCase := range testCases {
		// fmt.Println("fairCandySwap(", testCase[0], ", ", testCase[1], ") => ", fairCandySwap(testCase[0], testCase[1]))
		fmt.Println("fairCandySwapLinear(", testCase[0], ", ", testCase[1], ") => ", fairCandySwapLinear(testCase[0], testCase[1]))
	}
}

// per https://leetcode.com/articles/fair-candy-swap/
func fairCandySwapLinear(A []int, B []int) []int {
	// original total of Alice
	totalAlice := 0
    for i := 0; i < len(A); i++ {
		totalAlice += A[i]
	}

	// original total of Bob
	totalBob := 0
    for i := 0; i < len(B); i++ {
		totalBob += B[i]
	}

	result := []int{}
	delta := (totalBob - totalAlice) / 2	// If Alice gives x, she expects to receive x + delta
	bMap := map[int]int{}
	for _, val := range B {
		bMap[val] = val
	}

	for _, valA := range A {
		_, contains  := bMap[valA + delta]

		if contains {
			result = []int{valA, valA + delta}
		}
	}

	return result
}

// quadratic performance
func fairCandySwap(A []int, B []int) []int {
	// original total of Alice
	totalAlice := 0
    for i := 0; i < len(A); i++ {
		totalAlice += A[i]
	}

	// original total of Bob
	totalBob := 0
    for i := 0; i < len(B); i++ {
		totalBob += B[i]
	}

	// result placeholder
	result := []int{}

	swapA := 0	// candy value that Alice gives Bob
	swapB := 0	// candy value that Bob gives Alice
	for i := 0; i < len(A); i++ {
		swapA = A[i]

		for j := 0; j < len(B); j++ {
			swapB = B[j]

			// if Alice gives Bob swapA, and Bob gives Alice swapB, would their candy totals equal?
			if (totalAlice - swapA + swapB) == (totalBob - swapB + swapA) {
				// found a solution 
				result = []int{swapA, swapB}
			}
		}
	}

	return result
}
