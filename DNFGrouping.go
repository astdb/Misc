// "dutch national flag" regrouping
// given an int array and a pivot index, regroup elements into three successive groups
// in the array: elements lower than pivot, same as pivot and greater than the pivot

package main

import (
	"fmt"
	"errors"
)

func main () {
	testcase := []int{3,2,1,6,5,4,9,8,4,6,5,1,4,3,2,5,1,6,8,5,4,9,8,7,9,8,7,6,8,7,4,3,6,5,4}
	fmt.Println(DNFGrouping(testcase, 5))
}

func DNFGrouping(input []int, pivot int) ([]int, error) {
	if pivot >= len(input) {
		return nil, errors.New(fmt.Sprintf("DNFGrouping(): pivot (%d) outside valid range.", pivot))
	}

	// not in place - O(n) space
	res := []int{}
	low := []int{}
	equal := []int{}
	large := []int{}
	pivotVal := input[pivot]
	
	for _, v := range input {
		if v < pivotVal {
			low = append(low, v)
		}

		if v == pivotVal {
			equal = append(equal, v)
		}

		if v > pivotVal {
			large = append(large, v)
		}
	}

	fmt.Println("\tLow:", low)
	fmt.Println("\tEqual:", equal)
	fmt.Println("\tLarge:", large)

	res = append(res, low...)
	res = append(res, equal...)
	res = append(res, large...)

	return res, nil
}
