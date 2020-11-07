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
	freqs := map[int]int{}

	for _, n := range nums {
		_, ok := freqs[n]

		if ok {
			freqs[n]++
		} else {
			freqs[n] = 1
		}
	}

	freqMap := map[int]bool{}
	// dupFreq := false

	numFreqList := []*NumFreq{}
	for num, freq := range freqs {
		numFreqList = append(numFreqList, &NumFreq{Num: num, Freq: freq})

		_, ok := freqMap[freq]
		if ok {
			// repeated num frequency
			// dupFreq = true
		} else {
			freqMap[freq] = true
		}
	}

	sort.Slice(numFreqList, func(i, j int) bool {
		if numFreqList[i].Freq == numFreqList[j].Freq {
			return numFreqList[i].Num > numFreqList[j].Num
		}

		return numFreqList[i].Freq < numFreqList[j].Freq
	})

	res := []int{}
	for _, numFreq := range numFreqList {
		num := numFreq.Num
		freq := numFreq.Freq

		for i := 0; i < freq; i++ {
			res = append(res, num)
		}
	}

	return res
}

type NumFreq struct {
	Num  int
	Freq int
}
