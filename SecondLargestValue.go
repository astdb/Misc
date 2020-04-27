/*
Given an array of integers ,find the second largest value

Example 01:
	Input: [3,2,1,9,8,4,6,5,4]
	Output: 9

Example 02:
	Input: []
	Output:
*/

package main

import (
	"errors"
	"log"
)

func main() {
	tests := [][]int{{2, 7, 4, 1, 8, 1}, {1, 3}, {9, 3, 2, 10}, {0}, {}}

	for _, test := range tests {
		index, value, err := secLargest(test)
		if err != nil {
			log.Println(err)
			continue
		}

		log.Printf("secLargest(%v) == test[%d] == %d\n", test, index, value)
	}
}

// given an array of ints return the index and value of second largest element
func secLargest(x []int) (int, int, error) {
	var secLargestVal int
	var secLargestIndex int

	// find index of largest value
	var largestVal int
	var largestValIndex int
	for k, v := range x {
		if k == 0 {
			largestVal = v
			largestValIndex = k
		} else {
			if v > largestVal {
				largestVal = v
				largestValIndex = k
			}
		}
	}

	// log.Printf("\tsecLargest(): largestValIndex: %d\n", largestValIndex)

	if len(x) <= 0 {
		return secLargestIndex, secLargestVal, errors.New("secLargest(): empty input array.")
	}

	first := true
	for k, v := range x {
		// if v <= largestVal {
		if k != largestValIndex {
			// second largest value candidate
			// log.Printf("\tsecLargest(): considering x[%d] == %d as second largest candidate..\n", k, v)

			if first {
				// log.Printf("\tsecLargest(): initializing x[%d] == %d to sec largest val\n", k, v)
				secLargestVal = v
				secLargestIndex = k
				first = false
			} else {
				if v > secLargestVal {

					// log.Printf("\tsecLargest(): competing second largest value found at x[%d] == %d, resetting prev val..\n")

					secLargestVal = v
					secLargestIndex = k
				}
			}
		}
	}

	return secLargestIndex, secLargestVal, nil
}
