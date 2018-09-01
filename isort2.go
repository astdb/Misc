// Go function to perform, insertion sort on an input slice of ints

package main

import (
	"fmt"
)

func main() {
	// input := []int{5,3,4,6,1,2}
	input := []int{3,2}
	fmt.Println(isort(input))
}

func isort(a []int) []int {
	if len(a) <= 0 {
		return a
	}

	for i := 1; i < len(a); i++ {
		j := i-1
		curr := a[i]
		for j >= 0 && a[j] > a[i] {
			j--
		}

		k := i
		for k > j+1 {
			a[k] = a[k-1]
			k--
		}
		a[j+1] = curr
	}

	return a
}
