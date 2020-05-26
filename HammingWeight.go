package main

import (
	"log"
)

func main() {
	tests := []uint32{00000000000000000000000000001011, 00000000000000000000000010000000}

	for _, test := range tests {
		log.Printf("hammingWeight(%d) == %d\n", test, hammingWeight(test))
	}
}

func hammingWeight(num uint32) int {
hamWght := 0

	for num > 0 {
		if num % 10 == 1 {
		hamWght++
	}

	num = num / 10
	}

return hamWght
}

