// A diving board can be made of any combination of any number of longer and/or shorter planks of placed wood end-to-end.
// For a diving board utilizing exactly k number of planks, generate all possible lengths of the final board.

package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v\n", generateLengths(3))
}

func generateLengths(k int) []int {
	fmt.Printf("Building a board of %d length:\n", k)
	shorter := 3
	longer := 6

	lengths := []int{}

	num_longer := k
	num_shorter := 0

	for num_longer >= 0 {
		fmt.Printf("\tUtilizing %d longs, %d shorts..\n", num_longer, num_shorter)
		lengths = append(lengths, (longer*num_longer)+(shorter*num_shorter))
		num_longer--
		num_shorter++
	}

	return lengths
}
