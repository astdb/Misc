/*
Given an array arr of positive integers sorted in a strictly increasing order, and an integer k.

Find the kth positive integer that is missing from this array.

 

Example 1:

Input: arr = [2,3,4,7,11], k = 5
Output: 9
Explanation: The missing positive integers are [1,5,6,8,9,10,12,13,...]. The 5th missing positive integer is 9.
Example 2:

Input: arr = [1,2,3,4], k = 2
Output: 6
Explanation: The missing positive integers are [5,6,7,...]. The 2nd missing positive integer is 6.
 

Constraints:

1 <= arr.length <= 1000
1 <= arr[i] <= 1000
1 <= k <= 1000
arr[i] < arr[j] for 1 <= i < j <= arr.length

*/

package main

import (
  "log"
)

func main() {
 tests := [][][]int{{{2,3,4,7,11},{5}}, {{1,2,3,4},{2}}} 
  for _, test := range tests {
    log.Printf("findKthPositive(%v, %d) = %d\n", findKthPositive(test[0], test[1][0]), test[0], test[1][0])
  }
}

func findKthPositive(arr []int, k int) int {
  intCount := 1
  missingCount := 0
  arrIndex := 0

  for missingCount < k {
    if arrIndex < len(arr) {
      if arr[arrIndex] == intCount {
        // not a missing int - increment both int and array index
        intCount++
        arrIndex++
  
      } else {
        // current index value is bigger than current int - increase index until arr[index] == int (and count missing ints)
        if arr[arrIndex] > intCount {
            for intCount < arr[arrIndex] {
              missingCount++
              intCount++
  
              if missingCount == k {
                return intCount
              }
            }
  
            // at this point, arr[index] == int
            intCount++
            arrIndex++
        } else {
          // arr[index] < int : cannot happen?
        }
  
  
      }
    } else {
      for missingCount < k {
        missingCount++
        intCount++

        if missingCount == k {
           return intCount
        }
      }
    }   

    // intCount++
  }

  return intCount
}
