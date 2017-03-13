// given two arrays of numbers, compute the smallest non-negative difference between two numbers (one from each array)

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// testCases := []int{31, 16, 4, 2, 39, 20, 7, 10, 28, 18}
	// fmt.Printf("Minimum difference: %d\n", smallestDiffSingleArray(testCases))
	a1 := []int{1, 3, 15, 11, 2}
	a2 := []int{23, 127, 235, 19, 8}
	fmt.Println(smallestDiffDoubleArrays(a1, a2))
	return
}

// compute smallest diff in two arrays
func smallestDiffDoubleArrays(a1, a2 []int) int {
	sort.Ints(a1)
	sort.Ints(a2)

	a := 0 // pointer to array a1
	b := 0 // pointer to array a2
	var min int

	for i := 0; a < len(a1) && b < len(a2); i++ {
		if i == 0 {
			// initialize min
			min = a1[a] - a2[b]

			if min < 0 {
				min = min * (-1)
			}
		} else {
			if math.Abs(float64(a1[a]-a2[b])) < float64(min) {
				min = a1[a] - a2[b]

				if min < 0 {
					min = min * (-1)
				}
			}
		}

		if a1[a] < a2[b] {
			a++
		} else {
			b++
		}
	}

	return min
}

// compute smallest diff in single array
func smallestDiffSingleArray(numbers []int) int {
	// SORT ARRAY
	sort.Ints(numbers)
	fmt.Println(numbers)

	var diff int

	for i := 1; i < len(numbers); i++ {
		thisDiff := numbers[i] - numbers[i-1]

		if i == 1 || thisDiff < diff {
			diff = thisDiff
		}

		if diff == 0 {
			return diff
		}
	}

	return diff
}
