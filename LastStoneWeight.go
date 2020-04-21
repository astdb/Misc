/*
We have a collection of stones, each stone has a positive integer weight.

Each turn, we choose the two heaviest stones and smash them together.  Suppose the stones have weights x and y with x <= y.  The result of this smash is:

If x == y, both stones are totally destroyed;
If x != y, the stone of weight x is totally destroyed, and the stone of weight y has new weight y-x.
At the end, there is at most 1 stone left.  Return the weight of this stone (or 0 if there are no stones left.)



Example 1:

Input: [2,7,4,1,8,1]
Output: 1
Explanation:
We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of last stone.


Note:

1 <= stones.length <= 30
1 <= stones[i] <= 1000
*/

package main

import (
	"errors"
	"log"
)

func main() {
	tests := [][]int{{2, 7, 4, 1, 8, 1}, {1,3}}

	for _, test := range tests {
		log.Printf("lastStoneWeight(%v) == %d\n", test, lastStoneWeight(test))
	}
}

func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		log.Printf("==============================")
		log.Printf("Stones: %v\n", stones)
		// find largest  value
		largestIndex, largestVal, err := largest(stones)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Largest/index: %d/%d\n", largestVal, largestIndex)

		// find second largest value
		secLargestIndex, secLargestVal, err := secLargest(stones, largestVal)
		if err != nil {
			log.Fatal(err)
		}

		// smash result
		smashRes := largestVal - secLargestVal

		// create new stones array and copy all elements in except largest and second largest: copy smash result instead
		newStones := []int{}
		for k, v := range stones {
			if k != largestIndex && k != secLargestIndex {
				newStones = append(newStones, v)

			}
		}

		newStones = append(newStones, smashRes)
		stones = newStones

	}

	if len(stones) <= 0 {
		log.Fatalf("len(stones) <= 0")
	} 

	return stones[0]
}

// given an array of ints, and it's largest value, return the index and value of second largest element
func secLargest(x []int, largestVal int) (int, int, error) {
	var secLargestVal int
	var secLargestIndex int

	if len(x) <= 0 {
		return secLargestIndex, secLargestVal, errors.New("secLargest(): empty input array.")
	}

	for k, v := range x {
		first := true
		if v <= largestVal {
			// second largest value candidate
			if first {
				secLargestVal = v
				secLargestIndex = k
				first = false
			} else {
				if v > secLargestVal {
					secLargestVal = v
					secLargestIndex = k
				}
			}
		}
	}

	return secLargestIndex, secLargestVal, nil
}

// given an array of ints, return the index and value of largest element
func largest(x []int) (int, int, error) {
	var largestVal int
	var largestIndex int

	if len(x) <= 0 {
		return largestIndex, largestVal, errors.New("largest(): empty input array.")
	}

	for k, v := range x {
		if k == 0 {
			// initialize
			largestVal = v
			largestIndex = k
		} else {
			if v > largestVal {
				// update
				largestVal = v
				largestIndex = k
			}
		}
	}

	return largestIndex, largestVal, nil
}
