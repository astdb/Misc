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
  tests2 := [][]int{{0,2,1,-6,6,7,9,-1,2,0,1}, {0,2,1,-6,6,7,9,-1,2,0,1}, {3,3,6,5,-2,2,5,1,-9,4}}

  for _, test := range tests {
      log.Printf("TwoPartsEqSum(%v) == %v\n", test, TwoPartsEqSum(test))
  }

  for _, test := range tests2 {
      log.Printf("ThreePartsEqSum(%v) == %v\n", test, canThreePartsEqualSum(test))
  }
}

func canThreePartsEqualSum(A []int) bool {
  if len(A) < 3 {
    // cannot split A into three non-empty parts
    return false
  }

  aTotal := sumArr(A)
  if aTotal % 3 != 0 {
    // total of A's elements cannot be split into three
    return false
  }

  third := aTotal/3
  partSum := 0  // partial sum of A's elements up to i'th
  x1 := false   // flag indicating if there's a point in A where the sum of preceding elements == third
  x2 := false   // flag indicating if there's a point in A where the sum of preceding elements == (2*third)
  for i := 0; i < len(A); i++ {
    partSum += A[i]

    if partSum == third {
      x1 = true
    }

    if partSum == (2*third) {
      x2 = true
    }
  }

  if x1 && x2 {
    return true
  }

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
