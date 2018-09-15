// "dutch national flag" regrouping
// given an int array and a pivot index, regroup elements into three successive groups
// in the array: elements lower than pivot, same as pivot and greater than the pivot

package main

import (
	"fmt"
	"errors"
)

func main () {
	
}

func DNFGrouping(input []int, pivot int) ([]int, error) {
	if pivot >= len(input) {
		return nil, errors.New(fmt.Sprintf("DNFGrouping(): pivot (%d) outside valid range.", pivot))
	}

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

	res = append(res, low...)
	res = append(res, equal...)
	res = append(res, large...)
}
