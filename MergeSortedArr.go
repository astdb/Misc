/*
Given two sorted integier arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:

The number of elements initialized in nums1 and nums2 are m and n respectively.
You may assume that nums1 has enough space (size that is equal to m + n) to hold additional elements from nums2.
Example:

Input:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]
 

Constraints:

-10^9 <= nums1[i], nums2[i] <= 10^9
nums1.length == m + n
nums2.length == n
*/

package main

import (
	"log"
)

func main() {
	tests := [][][]int{ {{1,2,3,0,0,0}, {3}, {2,5,6}, {3}} }

	for _, test := range tests {
		log.Printf("merge(%v, %d, %v, %d) = %v\n", test[0], test[1][0], test[2], test[3][0], test[0])
	}
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
  nums1Index := 0
	for i := 0; i < len(nums2); i++ {
    for j := nums1Index; j < len(nums1); j++ {
      if nums1[j] > nums2[i] {
        // shift nums1 elements one up from index j, and insert nums2[i] at nums1[j]
        shift(nums2, j)
        nums1[j] = nums2[i]
        nums1Index = j+1
      }
    }
	}
}

func shift(arr []int, i int) {
  j := len(arr)-2

  for j >= i {
    arr[j+1] = arr[j]
    j--
  }
}

