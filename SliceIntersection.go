
/*
    Given two arrays, write a function to compute their intersection.

    Example:
    Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2].

    Note:
    Each element in the result must be unique.
    The result can be in any order.
*/

package main

import (
	"fmt"
)

func main() {
	v1 := []int{1,2,2,1}
	v2 := []int{2,2}

	fmt.Printf("%v ∩ %v == %v\n", v1, v2, intersection(v1, v2))
}

func intersection(nums1 []int, nums2 []int) []int {
	inter := []int{}

	for _,v1 := range nums1 {
		for _,v2 := range nums2 {
			if v1 == v2 && !contains(inter, v1) {
				inter = append(inter, v1)
			}
		}
	}

	return inter
}

func contains(nums []int, n int) bool {
	for _,v := range nums {
		if v == n {
			return true
		}
	}

	return false
}
