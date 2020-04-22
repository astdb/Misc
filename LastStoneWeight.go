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
	"sort"
)

func main() {
	tests := [][]int{{2, 7, 4, 1, 8, 1}, {1,3}, {9,3,2,10}}
	// tests := [][]int{{9,3,2,10}}

	for _, test := range tests {
		log.Printf("lastStoneWeight2(%v) == %d\n", test, lastStoneWeight2(test))
	}
}

func lastStoneWeight(stones []int) int {
	// sort stones by weight
	sort.Ints(stones)

	for len(stones) > 1 {
		// smash largest
		smash := stones[len(stones)-1] - stones[len(stones)-2]
		
		// copy over, replacing two largest with smash result
		newStones := []int{}
		for i := 0; i < (len(stones)-2); i++ {
			newStones = append(newStones, stones[i])
		}

		newStones = append(newStones, smash)
		stones = newStones
		sort.Ints(stones)
	}

	if len(stones) <= 0 {
		log.Fatal("len(stones) <= 0")
	}

	return stones[0]
}

func lastStoneWeight2(stones []int) int {
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
		secLargestIndex, secLargestVal, err := secLargest(stones, largestIndex)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Sec. largest/index: %d/%d\n", secLargestVal, secLargestIndex)

		// smash result
		smashRes := largestVal - secLargestVal
		log.Printf("SmashRes: %d\n", smashRes)

		// create new stones array and copy all elements in except largest and second largest: copy smash result instead
		newStones := []int{}
		for k, v := range stones {
			if k != largestIndex && k != secLargestIndex {
				newStones = append(newStones, v)

			}
		}

		newStones = append(newStones, smashRes)
		log.Printf("New Stones: %v\n", newStones)

		stones = newStones
		log.Printf("STONES: %v\n\n", stones)

	}

	if len(stones) <= 0 {
		log.Fatalf("len(stones) <= 0")
	} 

	return stones[0]
}

// given an array of ints, and it's largest value's index, return the index and value of second largest element
func secLargest(x []int, largestValIndex int) (int, int, error) {
	var secLargestVal int
	var secLargestIndex int

	log.Printf("\tsecLargest(): largestValIndex: %d\n", largestValIndex)

	if len(x) <= 0 {
		return secLargestIndex, secLargestVal, errors.New("secLargest(): empty input array.")
	}

	first := true
	for k, v := range x {
		// if v <= largestVal {
		if k != largestValIndex {
			// second largest value candidate
			log.Printf("\tsecLargest(): considering x[%d] == %d as second largest candidate..\n", k, v)
			
			if first {
				log.Printf("\tsecLargest(): initializing x[%d] == %d to sec largest val\n", k, v)
				secLargestVal = v
				secLargestIndex = k
				first = false
			} else {
				if v > secLargestVal {

					log.Printf("\tsecLargest(): competing second largest value found at x[%d] == %d, resetting prev val..\n")

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
