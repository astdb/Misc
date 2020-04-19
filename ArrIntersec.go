/*
Given two arrays, write a function to compute their intersection.

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]
Note:

Each element in the result should appear as many times as it shows in both arrays.
The result can be in any order.

Follow up:

What if the given array is already sorted? How would you optimize your algorithm?
What if nums1's size is small compared to nums2's size? Which algorithm is better?
What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?

https://leetcode.com/problems/intersection-of-two-arrays-ii/
*/

package main

import (
	"log"
)

func main() {
	tests := [][][]int{{{1, 2, 2, 1}, {2, 2}}, {{4, 9, 5}, {9, 4, 9, 8, 4}}}

	for _, test := range tests {
		// log.Printf("intersect2(%v, %v) == %v\n", test[0], test[1], intersect2(test[0], test[1]))
		log.Printf("intersect(%v, %v) == %v\n", test[0], test[1], intersect(test[0], test[1]))
	}
}

func intersect2(nums1 []int, nums2 []int) []int {
	map1 := map[int]int{}
	for _, i := range nums1 {
		_, ok := map1[i]
		if !ok {
			map1[i] = 1
		} else {
			map1[i]++
		}
	}

	map2 := map[int]int{}
	for _, i := range nums2 {
		_, ok := map2[i]
		if !ok {
			map2[i] = 1
		} else {
			map2[i]++
		}
	}

	res := []int{}
	for k, v := range map1 {
		v2, exists := map2[k]
		if exists {
			var commonNum int
			if v != v2 {
				commonNum = v - v2
				if commonNum < 0 {
					commonNum = (commonNum * (-1))
				}
			} else {
				commonNum = v
			}

			for i := 0; i < commonNum; i++ {
				res = append(res, k)
			}
		}
	}

	return res
}

func intersect(nums1 []int, nums2 []int) []int {
	// store second array elements and counts in map

	smaller := []int{}
	larger := []int{}

	if len(nums1) < len(nums2) {
		smaller = nums1
		larger = nums2
	} else {
		smaller = nums2
		larger = nums1
	}

	elemMap := map[int]int{}
	for _, elem := range smaller {
		_, exists := elemMap[elem]
		if !exists {
			elemMap[elem] = 1
		} else {
			elemMap[elem]++
		}
	}

	res := []int{}
	for _, elem := range larger {
		_, exists := elemMap[elem]

		if exists && (elemMap[elem] > 0) {
			// for i := 0; i < elemCount; i++ {
			// 	res = append(res, elem)
			// }

			elemMap[elem]--
			res = append(res, elem)
		}

	}

	return res
}
