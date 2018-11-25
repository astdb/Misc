/*
Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:

The number of elements initialized in nums1 and nums2 are m and n respectively.
You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.
Example:

Input:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]

*/

package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1,2,3,0,0,0}
	nums2 := []int{2,5,6}
	// merge(nums1, 3, nums2, 3)
	fmt.Println(merge(nums1, 3, nums2, 3))
}

func merge(nums1 []int, m int, nums2 []int, n int) []int  {
	// set cursors for the two arrays being merged
	n1 := 0;
	n2 := 0;

	for n2 < len(nums2) {
		if nums1[n1] > nums2[n2] {
			// shift nums1 element right by one position
			i := len(nums1)-2

			for i >= n1 {
				nums1[i+1] = nums1[i]
				i--
			}

			// insert nums2 value into nums1
			nums1[n1] = nums2[n2]
			n2++
		}

		n1++
		// n2++
	}

	return nums1
}
