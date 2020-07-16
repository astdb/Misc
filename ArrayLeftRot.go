package main

import (
	"log"
)

func main() {
	tests := [][]int32{{1, 2, 3, 4, 5}}

	for _, test := range tests {
		log.Printf("LeftRotate(%v) == %v\n", test, LeftRotate(test, 2))
	}
}

// left-rotate n, x times
func LeftRotate(n []int32, x int32) []int32 {
	if len(n) <= 1 {
		return n
	}

  var j int32
	for j = 0; j < x; j++ {
		tmp := n[len(n)-1]

		for i := len(n) - 1; i > 0; i-- {
			n[i] = n[i-1]
		}

		n[0] = tmp
	}

	return n
}
