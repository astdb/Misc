package main

import (
	"fmt"
)

func main() {
	tests := [][]int{{3, 2, 1, 9, 6, 8, 4, 6, 8, 4, 3, 5, 1, 8, 4}, {1}, {1, 2, 3}, {3, 2, 1}, {6, 5, 4, 9}, {3, 2, 1, 9, 8, 6, 3, 5}, {0}, {}, nil}
	for _, test := range tests {
		fmt.Println(test, mergeSort(test))
	}
}

func mergeSort(n []int) []int {
	if len(n) <= 1 {
		return n
	}

	return merge(mergeSort(n[0:len(n)/2]), mergeSort(n[len(n)/2:]))
}

func merge(n1, n2 []int) []int {
	res := []int{}
	i1 := 0
	i2 := 0

	for i1 < len(n1) && i2 < len(n2) {
		if n1[i1] < n2[i2] {
			res = append(res, n1[i1])
			i1++
		} else {
			res = append(res, n2[i2])
			i2++
		}
	}

	if i1 < len(n1) {
		res = append(res, n1[i1:]...)
	}

	if i2 < len(n2) {
		res = append(res, n2[i2:]...)
	}

	return res
}
