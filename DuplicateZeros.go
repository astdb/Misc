/*
Given a fixed length array arr of integers, duplicate each occurrence of zero, shifting the remaining elements to the right.

Note that elements beyond the length of the original array are not written.

Do the above modifications to the input array in place, do not return anything from your function.

Example 1:

Input: [1,0,2,3,0,4,5,0]
Output: null
Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]
Example 2:

Input: [1,2,3]
Output: null
Explanation: After calling your function, the input array is modified to: [1,2,3]
*/

package main

import (
	"log"
)

func main() {
	tests := [][]int{{1,0,2,3,0,4,5,0}, {1,2,3}}

	for testNo, test := range tests {
		log.Printf("#%d duplicateZeros(%v) == ", testNo, test)
		duplicateZeros(test)
		log.Printf("%v\n", test)
	}
}


func duplicateZeros(x []int)  {
	for i := 0; i < len(x); i++ {
		if x[i] == 0 {
			for j := len(x)-2; j > i; j-- {
				x[j]= x[j+1]
			}

			if i+1 < len(x) {
				x[i+1] = 0
				i++
			}
		}
	} 
}
