/*
Given an array A of integers, return true if and only if we can partition the array into three non-empty parts with equal sums.

Formally, we can partition the array if we can find indexes i+1 < j with (A[0] + A[1] + ... + A[i] == A[i+1] + A[i+2] + ... + A[j-1] == A[j] + A[j-1] + ... + A[A.length - 1])

 

Example 1:

Input: A = [0,2,1,-6,6,-7,9,1,2,0,1]
Output: true
Explanation: 0 + 2 + 1 = -6 + 6 - 7 + 9 + 1 = 2 + 0 + 1
Example 2:

Input: A = [0,2,1,-6,6,7,9,-1,2,0,1]
Output: false
Example 3:

Input: A = [3,3,6,5,-2,2,5,1,-9,4]
Output: true
Explanation: 3 + 3 = 6 = 5 - 2 + 2 + 5 + 1 - 9 + 4
 

Constraints:

3 <= A.length <= 50000
-10^4 <= A[i] <= 10^4
*/

package main

import (
  "log"
)

func main() {
  tests := [][]int{{3,3,2,6,2}, {0}, {0,0}, {1,2,3,4,5,6}, {1,2,3,9,4,5,6}}

  for _, test := range tests {
      log.Printf("TwoPartsEqSum(%v) == %v\n", test, TwoPartsEqSum(test))
  }
}

func canThreePartsEqualSum(A []int) bool {
  return false
}

// check if a given array can be split into two non-empty parts of equal sum
func TwoPartsEqSum(arr []int) bool {
  if len(arr) < 2 {
    // need at least two elements to partition array into two non-empty sections
    return false
  }

  for firstPartSize := 1; firstPartSize < len(arr); firstPartSize++ {
    if sumArr(arr[:firstPartSize]) == sumArr(arr[firstPartSize:]) {
      return true
    }
  }

  return false
}

// helper function to sum array
func sumArr(arr []int) int {
  sum := 0

  for _, x := range arr {
    sum += x
  }

  return sum
}
