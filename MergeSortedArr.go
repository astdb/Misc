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
	"fmt"
	"log"
)

func main() {
	// arr := []int{1,2,3,0,0,0}
	// log.Println(arr)
	// shift(arr, 2)
	// log.Println(arr)
	// return

	tests := [][][]int{ {{1, 2, 3, 0, 0, 0}, {3}, {2, 5, 6}, {3}}, {{0}, {0}, {1}, {1}}, {{2,0}, {1}, {1}, {1}} }

	for _, test := range tests {
		nums1 := test[0]
		m := test[1][0]
		nums2 := test[2]
		n := test[3][0]

		fmt.Printf("merge(%v, %d, %v, %d) = ", nums1, m, nums2, n)
		merge(nums1, m, nums2, n)
		fmt.Printf("%v\n", nums1)
	}
}

func merge(nums1 []int, m int, nums2 []int, n int) {
  nums1Index := 0
  
  if m == 0 {
    // whole of nums2 can just be appended to nums1
    for i := 0; i < len(nums1) && i < len(nums2); i++ {
      nums1[i] = nums2[i]
    }

    return
  }

	log.Printf("nums1: %v\n", nums1)
	log.Printf("nums2: %v\n", nums2)

	// for each element in the second array
	for i := 0; i < len(nums2); i++ {
		log.Printf("----------------merge(): i = %d\n", i)
		// find where it would fit in on first array

		log.Printf("merge(): starting nums1 scan from index %d\n", nums1Index)
		for j := nums1Index; j < len(nums1); j++ {
			// for j := nums1Index; nums1[j] > nums2[i]; j++ {
			log.Printf("merge(): j = %d\n", j)

			// if nums2[i] should fit in before nums1[j]
			if nums1[j] > nums2[i] {
				log.Printf("merge(): nums1[j (%d)] (%d) > nums2[i (%d)] (%d), looking to insert nums1[j (%d)] at nums2[i (%d)]\n", j, nums1[j], i, nums2[i], j, i)
				// shift nums1 elements one up from index j, and insert nums2[i] at nums1[j]
				shift(nums1, j)

				log.Printf("merge(): assigning nums2[i (%d)] (%d) to nums1[j (%d)] (%d)\n", i, nums2[i], j, nums1[j])
				nums1[j] = nums2[i]
				log.Printf("merge(): nums1: %v\n", nums1)

				nums1Index = j + 1
				m = m + 1
				log.Printf("merge(): incrementing nums1Index (now %d)\n", nums1Index)
				break
			}

			//
			if j == m-1 {
				// insert nums2[i:] from nums2[j+1]
				x := j + 1
				y := i
				for {
					nums1[x] = nums2[y]
					x++
					y++

					if x > len(nums1)-1 || y > len(nums2)-1 {
						// break
						return
					}
				}
			}
		}
	}
}

func shift(arr []int, i int) {
	log.Println("shift(): shifting..")

	log.Printf("shift(): pre-shift: %v\n", arr)
	j := len(arr) - 2

	for j >= i {
		arr[j+1] = arr[j]
		j--
	}

	log.Printf("shift(): post-shift: %v\n", arr)
}
