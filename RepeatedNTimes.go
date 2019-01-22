
/*
In a array A of size 2N, there are N+1 unique elements, and exactly one of these elements is repeated N times.

Return the element repeated N times.

 

Example 1:

Input: [1,2,3,3]
Output: 3
Example 2:

Input: [2,1,2,5,3,2]
Output: 2
Example 3:

Input: [5,1,5,2,5,3,5,4]
Output: 5
 

Note:

4 <= A.length <= 10000
0 <= A[i] < 10000
A.length is even
*/

package main

import (
	"fmt"
)

func main() {
	tests := [][]int{{1,2,3,3}, {2,1,2,5,3,2}, {5,1,5,2,5,3,5,4}}

	for _, test := range tests {
		fmt.Println(repeatedNTimes(test))
	}
}

func repeatedNTimes(nums []int) int {
	n := len(nums)/2
	numsElemCounts := map[int]int{}

	for _, elem := range nums {
		_, exists := numsElemCounts[elem]
		if exists {
			numsElemCounts[elem]++
		} else {
			numsElemCounts[elem] = 1
		}

		if numsElemCounts[elem] == n {
			return elem
		}
	}

	return 0
}
