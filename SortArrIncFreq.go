/*
Given an array of integers nums, sort the array in increasing order based on the frequency of the values. If multiple values have the same frequency, sort them in decreasing order.

Return the sorted array.


Example 1:

Input: nums = [1,1,2,2,2,3]
Output: [3,1,1,2,2,2]
Explanation: '3' has a frequency of 1, '1' has a frequency of 2, and '2' has a frequency of 3.

Example 2:

Input: nums = [2,3,1,3,2]
Output: [1,3,3,2,2]
Explanation: '2' and '3' both have a frequency of 2, so they are sorted in decreasing order.
Example 3:

Input: nums = [-1,1,-6,4,5,-6,1,4,1]
Output: [5,-1,4,4,-6,-6,1,1,1]

*/

package main

import (
  "log"
  "sort"
)

func main() {
	tests := [][]int{{1, 1, 2, 2, 2, 3}, {2, 3, 1, 3, 2}}

	for _, test := range tests {
		log.Printf("frequencySort(%v) = %v\n", test, frequencySort(test))
	}
}

func frequencySort(nums []int) []int {
	// build map of nums element frequencies
	freqs := map[int]int{}
	keys := []int{} // make a list of keys

	for _, num := range nums {
		_, counted := freqs[num]
		if counted {
			freqs[num]++
		} else {
			freqs[num] = 1
			keys = append(keys, num)
		}
	}

	if multValsInMap(freqs) {
		sort.Ints(keys)
	} else {
		// sort.Reverse(sort.Ints(keys))
    sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	}

	res := []int{}

	for _, k := range keys {
		keyNum := freqs[k]

		for i := 0; i < keyNum; i++ {
			res = append(res, k)
		}
	}

	return keys
}

// check if a given int->int map has multiples of the same value
func multValsInMap(mp map[int]int) bool {
	// interim map keyed by values in mp
	valuesMap := map[int]bool{}
	for _, v := range mp {
		_, present := valuesMap[v]
		if present {
			// v has been seen before
			return true
		} else {
			valuesMap[v] = true
		}
	}

	return false
}
