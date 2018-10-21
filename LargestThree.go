// given a list of unique positive numbers, find the largest three in linear time.

package main

import (
	"fmt"
)

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(input)
	fmt.Println(LargestThree(input))
}

func LargestThree(list []int) []int {
	result := []int{} // placeholder to hold three largest elements in list
	if len(list) <= 0 {
		// return empty list if input is empty
		return result
	}

	// search through the list thrice, to find three largest elements
	for k := 0; k < 3; k++ {
		var largest int // placeholder for the largest element
		j := 0          // custom loop counter
		for i := 0; i < len(list); i++ {
			// if this element is not counted as a large element yet
			if !IsIn(list[i], result) {
				if j == 0 {
					// initialize largest placeholder, if this is the first element considered
					largest = list[i]
				} else {
					// if not the first element considered, check if it's larger than the currently know largest element (for this iteration)
					if list[i] > largest {
						largest = list[i]
					}
				}

				j++ // increment custom counter
			}
		}

		result = append(result, largest) // add largest found in this iteration to result array
	}

	return result
}

// checks if a given number is in a given list
func IsIn(i int, list []int) bool {
	for _, v := range list {
		if i == v {
			return true
		}
	}

	return false
}
